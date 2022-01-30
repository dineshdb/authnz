package auth

import (
	"github.com/golang-jwt/jwt"
)

type Token struct {
	AccessToken  string `json:"accessToken" validate:"nonzero"`
	RefreshToken string `json:"refreshToken" validate:"nonzero"`
}

type TokenClaims struct {
	jwt.StandardClaims
	Scope string
}
