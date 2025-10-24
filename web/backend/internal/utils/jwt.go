package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

const SecretKey = "flexgrid_secret"

func GenerateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
	})

	return token.SignedString([]byte(SecretKey))
}

func ParseToken(tokenStr string) (*int, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("incorrect signature algorithm")
		}

		return []byte(SecretKey), nil
	}

	token, err := jwt.Parse(tokenStr, keyFunc)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		floatUserId := claims["user_id"].(float64)

		userId := int(floatUserId)

		return &userId, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, nil
}
