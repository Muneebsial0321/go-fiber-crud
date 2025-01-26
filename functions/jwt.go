package functions

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var key = []byte("MY_KEY")

func Jwt_SignToken(payload string) (string, error) {
	claims := jwt.MapClaims{
		"userId": payload,
	}
	metadata := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := metadata.SignedString(key)
	if err != nil {
		return "", err
	}
	return token, nil
}

func Jwt_VerifyToken(token string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return key, nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}
