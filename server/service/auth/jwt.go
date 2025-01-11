package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lb-developer/jotjournal/config"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	if string(secret) == "" {
		return "", fmt.Errorf("secret was empty when creating JWT")
	}

	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"issuedAt":  time.Now().Unix(),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
