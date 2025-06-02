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

	expiration := time.Second * time.Duration(config.Envs.RefreshExpirationInSeconds)
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

func ProtectedRoute(store types.UserStore, sessionStore types.SessionStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			sessionToken := getTokenFromRequest(req)
			if sessionToken == "" {
				log.Printf("sessionToken was empty")
				utils.WriteJSON(w, http.StatusUnauthorized, fmt.Errorf("sessionToken was empty"))
				return
			}

			// check if user has a current valid session
			_, err := sessionStore.ValidateSessionToken(sessionToken)
			if err != nil {
				log.Printf("Couldn't validate user session: %v", err)
				utils.WriteJSON(w, http.StatusUnauthorized, err)
				return
			}

			// check token from cache is valid
			token, err := ValidateToken(sessionToken)
			if err != nil {
				log.Printf("Couldn't validate JWT token: %v", err)
				utils.WriteJSON(w, http.StatusUnauthorized, err)
				return
			}

			// extract userID from token
			claims := token.Claims.(jwt.MapClaims)
			userIDString := claims["userID"].(string)

			userID, _ := strconv.Atoi(userIDString)
			user, err := store.GetUserByID(userID)
			if err != nil {
				log.Printf("Couldn't get user, error: %v", err)
				utils.WriteJSON(w, http.StatusUnauthorized, err)
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

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		claims := t.Claims.(jwt.MapClaims)
		// has the token expired
		if expFloat, ok := claims["expiresAt"].(float64); ok {
			expiresAt := int64(expFloat)
			if time.Now().Unix() > expiresAt {
				return token, fmt.Errorf("Token expired")
			}
		}

		return []byte(config.Envs.SessionSecret), nil
	})
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, found := ctx.Value(UserKey).(int)
	if !found {
		return -1
	}

	return userID
}
