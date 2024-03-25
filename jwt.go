package main

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtCustomClaims struct {
	UserId int64 `json:"userId"`
	jwt.RegisteredClaims
}

func makeToken(userId int64) (string, error) {
	claims := &jwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("zhuzhu"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
