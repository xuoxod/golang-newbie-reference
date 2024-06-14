package utils

import (
	"errors"
	"fmt"
	"time"
	"xuoxod/adminhelper/internal/constants"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(tokenString string) (*jwt.Token, bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SecretKey), nil
	})

	if err != nil {
		return nil, false, err
	}

	return token, token.Valid, nil
}

func GenerateJwt(user interface{}) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    fmt.Sprintf("%v", user),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(constants.SecretKey))

	if err != nil {
		return "", errors.New("Unabled to generate token")
	}

	return token, nil
}

func GenerateJwtFor(user interface{}, timeFrame int) (string, error) {
	var tf time.Duration = time.Duration(timeFrame)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    fmt.Sprintf("%v", user),
		ExpiresAt: time.Now().Add(time.Hour * tf).Unix(),
	})

	token, err := claims.SignedString([]byte(constants.SecretKey))

	if err != nil {
		return "", errors.New("Unabled to generate token")
	}

	return token, nil
}
