package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lb-developer/jotjournal/config"
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

func TestJWTExpirationTime(t *testing.T) {
	secret := []byte("valid-secret")
	tokenStr, _ := CreateJWT(secret, 1)

	token, _ := jwt.Parse(tokenStr, nil)

	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		t.Errorf("Couldn't get JWT claims")
	} else {
		issuedAtClaim, exists := claims["issuedAt"]
		if !exists {
			t.Errorf("claims did not contain field 'issuedAt'")
		}

		expiresAtClaim, exists := claims["expiresAt"]
		if !exists {
			t.Errorf("claims did not contain field 'expiresAt'")
		}

		issuedAt, convert := issuedAtClaim.(float64)
		if !convert {
			t.Errorf("Couldn't convert issuedAt to float64")
		}

		expiresAt, converted := expiresAtClaim.(float64)
		if !converted {
			t.Errorf("Couldn't convert expiresAt to float64")
		}

		issued := time.Duration(issuedAt * float64(time.Second))
		expires := time.Duration(expiresAt * float64(time.Second))
		minutes := (expires - issued).Minutes()
		expirationInMinutes := (time.Second * time.Duration(time.Duration(config.Envs.RefreshExpirationInSeconds))).Minutes()

		if minutes != expirationInMinutes {
			t.Errorf("Issued JWT Token expiration was %v minutes not 10 minutes", minutes)
		}
	}
}
