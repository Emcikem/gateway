package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"time"
)

var MySigningKey = []byte("asfasfdafasdfdasfa.")

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": username,
		"exp":  time.Now().Unix() + 1,
	})
	tokenString, err := token.SignedString(MySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetUsername(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return MySigningKey, nil
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprint(token.Claims.(jwt.MapClaims)["name"]), nil
}
