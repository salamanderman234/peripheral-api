package entity

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
