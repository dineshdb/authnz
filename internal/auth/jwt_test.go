package auth

import (
	"testing"

	"github.com/dineshdb/authnz/internal/user"
)

var pem string = `
-----BEGIN PRIVATE KEY-----
MIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQD2b/VXADM7r08b
xASKMKUFqGLbycdSQ7I82CdeNz0VlKc1sp1mYzk9B1czKFjUknc0guy+dXWt/0RR
DDiKcXNzdmDmCvEli3C1BagtYvctJ+DaYpWQhpIEJGojaNKO97CK0Id3xNDl/e5h
EExbri8IJ86SGnr4ubpuriXk0i9QzkbBcVh0S9xUZuB7U7i1ijsGVbn7/xi/18no
vd1aR84WHNQ4tNgllx27aPBzA6Ec2VqmWSjHkDo9A8SmBRePRSPy82p6wO8ib9pS
tIEGR6BNzv1RIZUR4ObS1jDESof08smRvnWpoAHudS1gaXKFnfsFeHucE9fKf/2C
LNfxOF6vAgMBAAECggEBAM+a3cofJwn+09wGM/TeqgasJiwWPk41LXBIgFHEozcM
9hgskqDwsgWRq4ozUTIy+S1JpnuEpFCinUDR1Mf8b1Azx8nEKgaBA7/cNiOWHbjy
wV/4cRtB4ryOmMOfyNIcI6OtrJHfQkSeuTUX79vET2bFciZvHG1wuXgISXANCUM9
++87/yqB2pSDwwLFnvaCTP7XeDTmBOfbd6R9TS6slLXpcbh5/Y/i3nOEJkHlaafl
mTgcXGw5itSH2/1qayNbtbI9sjtuEQQucf+B43UmkSeUeKjgJO6j6V7MsUssnkBt
dP8/3SGpj4Z+R1/VRNRDKfDsWTbXn0vv4pIWOfySFXECgYEA/kdU/QNPUrjGAUc4
beO3HhdF0HCvXic4FpDfrkvKyx+P9lews1y9gF2M5ghbAkzty3ueSX0J15qW1Vaw
BJRWQXRgyFFhbKJwIK9d6zkRAQglYcfIN1Nrf67xrPpOaFic03IcAnZ0kN/JsRye
6BnnMpKT++9b9CuF91wrK+cG+F0CgYEA+BsJjF3xYFHdK+wX8hkDSscSZuwzquo7
Al1Oo1SD1wDogs1w1yCYe+yISCa9YH8J3qcqix7WN7AN/QnOPdNeIH+dP+bXoTQ5
5euLK/ctOfv5v+bS9q6RocY/J8RRrQUtFfA0eL0MLw5m3rt0ro/+H98T+FabuW2D
vsAhpbiCknsCgYEAj7xWyGb0kfgsxVAzD7snKfVR24+3MevNgsQGDQp+6e8/e6r6
EYmc/VDkcqvKdjRyPxHz2eq6g4u5M4M7IHuRfpKAmvuVrMjtxSwcVPj/KawnJWy/
OrcHDzgfGP6tD8L3c3cPajz3i2VVJ67cDKuHy0icKk+VlSJ9KeSJ6tk/UWkCgYEA
2eU89IcI1xPuj4WQ3jFzf7foBHZLRj7iRkhWKQGvrCMDEOWGxZi98pAgfGVxio7n
xyC/L2GMt2mqT2HOPOQmVZpeK2H8XHp2ouPD3X/+u179z7jT7IwSIKbwjmdPaAoU
t6C3JJa7XZRjahft+OVDRRBxBHhj2W1B+EPbCSVLn4MCgYEA02Lxd/yd2XkY9K4j
/7ImIW6L/LZblTNxI7o/U2iI8jO7EzZ33yzXpxTDmtbEZox+6/6w2P7mceGb4IqW
Nf6ZKurwbOSIrU0HpuevoQkCk4NIQK5asVePF93pnSpc1JXqRjvFsZk57nen7ZSq
+aspV99yZkZ/e/EZGmdfcj2i7PQ=
-----END PRIVATE KEY-----
`

var jwtValidator JWTValidator = WithKey([]byte(pem))

func TestAccessToken(t *testing.T) {
	token, err := jwtValidator.Generate(user.User{ID: 1})
	if err != nil {
		t.Error("Token verification failed")
	}
	tokenClaims, verificationError := jwtValidator.Verify(token.AccessToken)
	if verificationError != nil {
		t.Error("Token verification failed")
	}

	if tokenClaims.Subject != "1" {
		t.Error("Token Subject is Different")
	}
}

func TestRefreshToken(t *testing.T) {
	token, err := jwtValidator.Generate(user.User{ID: 1})
	if err != nil {
		t.Error("Token verification failed")
	}
	newToken, err := jwtValidator.RefreshToken(token.RefreshToken)
	if err != nil {
		t.Error("Token refresh failed")
	}

	tokenClaims, verificationError := jwtValidator.Verify(newToken.AccessToken)
	if verificationError != nil {
		t.Error("Token verification failed")
	}

	if tokenClaims.Subject != "1" {
		t.Error("Token Subject is Different")
	}
}
