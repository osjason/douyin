package model

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestJWTCreate(t *testing.T) {
	claim, _ := NewJWTClaims(0)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte("key"))
	if err != nil {
		log.Fatalln("生成密钥失败")
	}
	t.Logf("JWT Token : %v", ss)
}

func TestStoreTokenByRedis(t *testing.T) {
	user1 := &User{
		Id:            0,
		Name:          "jason",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	claims, _ := NewJWTClaims(user1.Id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("key"))
	if err != nil {
		log.Fatalln("生成密钥失败")
	}
	t.Logf("JWT Token : %v", ss)

	// ---

	ctx := context.Background()
	rdc := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "admin",
		DB:       0,
	})

	if rdc == nil {
		t.Fatal("redis 客户端创建失败")
	}
	stopDuration, _ := time.ParseDuration("30s")
	err = rdc.Set(ctx, strconv.Itoa(user1.Id), ss, stopDuration).Err()
	if err != nil {
		t.Fatal("Redis Set Wrong")
	}

	time.Sleep(25 * time.Second)
	result, err := rdc.Get(ctx, strconv.Itoa(user1.Id)).Result()
	if err != nil {
		t.Fatal("25 秒后，redis Get 出错")
	}
	t.Logf("25秒时，尝试获取 Redis 键值，成功：%v", result)

	time.Sleep(6 * time.Second)

	result, err = rdc.Get(ctx, strconv.Itoa(user1.Id)).Result()
	if err != nil {
		t.Fatal("30秒后，redis Get 出错")
	}
	t.Log("30秒后，Redis get 仍然成功")
}
