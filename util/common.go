package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var MySigningKey = []byte("asfasfdafasdfdasfa.")

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

func InStringSlice(slice []string, str string) bool {
	for _, item := range slice {
		if str == item {
			return true
		}
	}
	return false
}
