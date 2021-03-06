package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username,omitempty"`
	Rules    uint16 `json:"rules,omitempty"`
}
