package model

import "github.com/dgrijalva/jwt-go"

type UserParams struct {
	NickName string `json:"Nick"`
	Admin bool `json:"Admin"`
	Password string `json:"Password"`
	Address string `json:"Address"`
	jwt.StandardClaims
}
