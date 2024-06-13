package model

import "github.com/golang-jwt/jwt"

type Auth struct {
	Token        string
	RefreshToken string
}

type AcessTokenClaims struct {
	Sub  int    `json:"sub"`
	Name string `json:"name"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	Sub int `json:"sub"`
	jwt.StandardClaims
}
