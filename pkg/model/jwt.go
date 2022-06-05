package model

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTClaims struct {
	UserId int
	jwt.StandardClaims
}

const (
	JWTTokenDuration time.Duration = 1 * time.Minute
	JWTSignMethod                  = "HS256"
)

func NewJWTClaims(user *User) (jwt.Claims, error) {
	issueTime := time.Now()

	return JWTClaims{
		UserId: (*user).Id,
		StandardClaims: jwt.StandardClaims{
			Audience:  "douyin-app",
			Issuer:    "douyin-api",
			IssuedAt:  issueTime.Unix(),
			NotBefore: issueTime.Unix(),
			ExpiresAt: issueTime.Add(JWTTokenDuration).Unix(),
		},
	}, nil
}
