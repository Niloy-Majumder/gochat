package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gochat/types/constants"
	"time"
)

func CreateJWTToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(constants.JWT_SECRET))
}

func VerifyJWTToken(jwtToken string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fiber.NewError(401, "invalid token")
	}

	return claims, nil
}
