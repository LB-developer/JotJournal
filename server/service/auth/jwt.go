package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lb-developer/jotjournal/config"
	"github.com/lb-developer/jotjournal/types"
	"github.com/lb-developer/jotjournal/utils"
)

type contextKey string

const UserKey contextKey = "userID"

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

func ProtectedRoute(store types.UserStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			tokenString := getTokenFromRequest(req)
			if tokenString == "" {
				log.Printf("Token string was empty")
				utils.WriteJSON(w, http.StatusUnauthorized, fmt.Errorf("Token string was empty"))
				return
			}

			token, err := validateToken(tokenString)
			if err != nil {
				log.Printf("Couldn't validate JWT token: %v", err)
				utils.WriteJSON(w, http.StatusForbidden, err)
				return
			}

			claims := token.Claims.(jwt.MapClaims)
			userIDString := claims["userID"].(string)

			userID, _ := strconv.Atoi(userIDString)
			user, err := store.GetUserByID(userID)
			if err != nil {
				log.Printf("Couldn't get user, error: %v", err)
				utils.WriteJSON(w, http.StatusForbidden, err)
				return
			}
			// inject userID into context
			ctx := req.Context()
			ctx = context.WithValue(ctx, UserKey, user.ID)
			req = req.WithContext(ctx)

			next.ServeHTTP(w, req)
		})
	}
}

func getTokenFromRequest(req *http.Request) string {
	return req.Header.Get("Authorization")
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, found := ctx.Value(UserKey).(int)
	if !found {
		return -1
	}

	return userID
}
