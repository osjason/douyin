package model

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// JWTClaims 做为 jwt.claims 类型参数
type JWTClaims struct {
	UserId int
	jwt.StandardClaims
}

const (
	JWTTokenDuration time.Duration = 1 * time.Minute
	JWTSignMethod                  = "HS256"
)

func NewJWTClaims(userId int) (JWTClaims, error) {
	issueTime := time.Now()

	return JWTClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Audience:  "douyin-app",
			Issuer:    "douyin-auth",
			IssuedAt:  issueTime.Unix(),
			NotBefore: issueTime.Unix(),
			// ExpiresAt: 日期由 redis 的过期功能定
		},
	}, nil
}
