package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserId uint `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uint) (string, error) {
	var key = []byte(os.Getenv("JWT_KEY"))
	claims := JWTClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "yubin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(key)
	return ss, err
}

func ValidateToken(tokenString string) (*JWTClaims, error) {
	var key = []byte(os.Getenv("JWT_KEY"))
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
