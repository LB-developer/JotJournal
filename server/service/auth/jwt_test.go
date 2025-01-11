package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestValidJWTCreation(t *testing.T) {
	secret := []byte("valid-secret")
	tokenStr, err := CreateJWT(secret, 1)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if tokenStr == "" {
		t.Errorf("token was empty")
	}
}

func TestEmptyJWTCreation(t *testing.T) {
	secret := []byte("")
	_, err := CreateJWT(secret, 1)

	if err == nil {
		t.Errorf("expected error for an empty secret")
	}
}

func TestJWTSigningMethod(t *testing.T) {
	secret := []byte("valid-secret")
	tokenStr, _ := CreateJWT(secret, 1)

	jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			t.Errorf("Expected signing method 'HMAC' got '%v'", token.Header["alg"])
		}
		return nil, nil
	})
}

