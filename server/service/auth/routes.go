package auth

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lb-developer/jotjournal/config"
	"github.com/lb-developer/jotjournal/types"
	"github.com/lb-developer/jotjournal/utils"
)

type Handler struct {
	sessionStore types.SessionStore
}

func NewHandler(sessionStore types.SessionStore) *Handler {
	return &Handler{
		sessionStore: sessionStore,
	}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Get("/refresh", h.handleRefreshToken)
}

// @Summary Renews access tokens
// @Description Validates the users refresh token and if valid returns a new access token
// @Tags Auth
// @Accepts json
// @Produce json
// @Param RefreshToken body types.RefreshTokenPayload true "refresh token"
// @Success 200 {object} types.AccessTokenResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/auth/refresh [post]
func (h *Handler) handleRefreshToken(w http.ResponseWriter, req *http.Request) {
	oldSessionToken := req.Header.Get("Authorization")

	// get userID from token claims
	token, _ := ValidateToken(oldSessionToken)
	claims := token.Claims.(jwt.MapClaims)
	userIDString := claims["userID"].(string)
	userID, _ := strconv.Atoi(userIDString)

	// is long-term session in database
	valid, err := h.sessionStore.ValidateSession(int64(userID), oldSessionToken)
	if !valid {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	// get sessionID to set in cache
	sessionID, err := h.sessionStore.ValidateSessionToken(oldSessionToken)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// create new short session token
	newToken, err := CreateJWT([]byte(config.Envs.SessionSecret), userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// clear old short session
	h.sessionStore.ClearSessionFromCache(oldSessionToken)

	// cache new short session
	_, err = h.sessionStore.CacheSessionToken(newToken, sessionID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// return new access token
	utils.WriteJSON(w, http.StatusOK, types.SessionTokenResponse{SessionToken: newToken})
}
