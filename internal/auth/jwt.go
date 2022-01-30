package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/dineshdb/authnz/internal/user"
	"github.com/golang-jwt/jwt"
)

type JWTValidator struct {
	privateKey                 *rsa.PrivateKey
	Issuer                     string
	AccessTokenExpiryDuration  time.Duration
	RefreshTokenExpiryDuration time.Duration
}

func New(privateKeyPath string) JWTValidator {
	privateKeyBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}
	return WithKey(privateKeyBytes)
}

func WithKey(privateKeyBytes []byte) JWTValidator {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	return JWTValidator{
		privateKey:                 privateKey,
		Issuer:                     "ACME Corp",
		AccessTokenExpiryDuration:  1 * time.Hour,
		RefreshTokenExpiryDuration: 1 * time.Hour,
	}
}

func (j *JWTValidator) Generate(user user.User) (Token, error) {
	tokenClaims := TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.AccessTokenExpiryDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    j.Issuer,
			// We use the user id instead of email to prevent leakage of email
			Subject: strconv.Itoa(int(user.ID)),
		},
	}

	accessTokenClaims := tokenClaims
	accessTokenClaims.Scope = "profile"
	accessToken, err := j.getToken(accessTokenClaims)
	if err != nil {
		return Token{}, err
	}

	// Use opaque token for refresh token
	refreshTokenClaims := tokenClaims
	refreshTokenClaims.Scope = "refresh"
	marshalled, err := json.Marshal(refreshTokenClaims)
	if err != nil {
		return Token{}, err
	}

	opaqueToken, err := rsa.EncryptPKCS1v15(rand.Reader, &j.privateKey.PublicKey, marshalled)
	if err != nil {
		return Token{}, err
	}

	return Token{
		AccessToken:  accessToken,
		RefreshToken: base64.StdEncoding.EncodeToString(opaqueToken),
	}, nil
}

func (j *JWTValidator) Verify(signedJWT string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(signedJWT, &TokenClaims{}, func(jwt *jwt.Token) (interface{}, error) {
		return &j.privateKey.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func (j *JWTValidator) VerifyRefreshToken(token string) (int, error) {
	raw, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, fmt.Errorf("bad base64 input")
	}

	decoded, err := rsa.DecryptPKCS1v15(rand.Reader, j.privateKey, raw)
	if err != nil {
		return 0, fmt.Errorf("badly encrypted token")
	}

	var tokenClaims TokenClaims
	if err := json.Unmarshal(decoded, &tokenClaims); err != nil {
		return 0, fmt.Errorf("invalid token body")
	}

	if !strings.Contains(tokenClaims.Scope, "refresh") {
		return 0, fmt.Errorf("not a refresh token")
	}

	id, err := strconv.Atoi(tokenClaims.Subject)
	if err != nil {
		return 0, fmt.Errorf("invalid user id")
	}
	return id, nil
}

func (j *JWTValidator) RefreshToken(token string) (Token, error) {
	id, err := j.VerifyRefreshToken(token)
	if err != nil {
		return Token{}, fmt.Errorf("invalid refresh token")
	}

	newToken, err := j.Generate(user.User{ID: int64(id)})
	if err != nil {
		return Token{}, fmt.Errorf("error generating token")
	}

	return newToken, nil
}

func (j *JWTValidator) getToken(claims TokenClaims) (string, error) {
	signer := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	return signer.SignedString(j.privateKey)
}
