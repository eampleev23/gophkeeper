package main

import "github.com/golang-jwt/jwt/v4"

// Claims описывает утверждения, хранящиеся в токене + добавляет кастомное UserID.
type Claims struct {
	jwt.RegisteredClaims
	UserID    int
	UserLogin string
}
