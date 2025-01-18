package config

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWTKEY"))

func CreateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  userId,
			"exp": time.Now().Add(time.Hour * 48).Unix(),
		})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (int, error) {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, err
	}

	return int(claims["id"].(float64)), err
}

func CheckAuthorized(tokenString string) (int, error) {
	if tokenString == "" {
		return 0, fmt.Errorf("missing authorization header")
	}

	id, err := VerifyToken(tokenString)

	if err != nil {
		return 0, err
	}

	return id, nil
}
