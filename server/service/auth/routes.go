package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lb-developer/jotjournal/config"
	"github.com/lb-developer/jotjournal/types"
	"github.com/lb-developer/jotjournal/utils"
)

type Handler struct {
	userStore types.UserStore
	authStore types.SessionStore
}

func NewHandler(authStore types.SessionStore) *Handler {
	return &Handler{
		authStore: authStore,
	}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Group(func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Get("/refresh", h.handleRefreshToken)
		})
	})
}

// @Summary Renews access tokens
// @Description Validates the users refresh token and if valid returns a new access token
// @Tags Auth
// @Accepts json
// @Produce json
// @Param RefreshToken body types.RefreshTokenPayload true "refresh token"
// @Success 200
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/auth/refresh [post]
func (h *Handler) handleRefreshToken(w http.ResponseWriter, req *http.Request) {
	userID := GetUserIDFromContext(req.Context())

	var refreshToken string 
	err := utils.ParseJSON(req, &refreshToken)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	valid, err := h.authStore.ValidateSession(int64(userID), refreshToken)
	if !valid {
		utils.WriteError(w, http.StatusUnauthorized, err)	
		return
	}

	// the refresh token is valid
	token, err := CreateJWT([]byte(config.Envs.JWTSecret), userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)	
		return
	}

	// return new access token
	utils.WriteJSON(w, http.StatusOK, types.AccessTokenResponse{ AccessToken: token})
}
