package jots

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lb-developer/jotjournal/service/auth"
	"github.com/lb-developer/jotjournal/types"
	"github.com/lb-developer/jotjournal/utils"
)

type Handler struct {
	jotStore     types.JotStore
	userStore    types.UserStore
	sessionStore types.SessionStore
}

func NewHandler(jotStore types.JotStore, userStore types.UserStore, sessionStore types.SessionStore) *Handler {
	return &Handler{
		jotStore:     jotStore,
		userStore:    userStore,
		sessionStore: sessionStore,
	}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Group(func(r chi.Router) {
		r.Use(auth.ProtectedRoute(h.userStore, h.sessionStore))

		r.Route("/jots", func(r chi.Router) {
			r.Get("/", h.handleGetJotsByUserID)
			r.Patch("/", h.handleUpdateJotByJotID)
		})
	})
}

// @Summary Get jots for the authenticated user
// @Description Retrieves all jots associated with the authenticated user based on their ID for the given month
// @Tags jots
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "JWT access token for authentication"
// @Param month query string true "jot search by month"
// @Success 200 {array} types.Jots
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 403 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/jots [get]
func (h *Handler) handleGetJotsByUserID(w http.ResponseWriter, req *http.Request) {
	userID := auth.GetUserIDFromContext(req.Context())
	if userID == -1 {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("couldn't find user id"))
		return
	}

	month := req.URL.Query().Get("month")
	if month == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no month in request"))
		return
	}

	intMonth, err := strconv.Atoi(month)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	jots, err := h.jotStore.GetJotsByUserID(intMonth, int64(userID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, jots)
}

// @Summary Updates jot for the authenticated user
// @Description Updates a jot associated with the authenticated user based on the jot ID
// @Tags jots
// @Security BearerAuth
// @Param Authorization header string true "JWT access token for authentication"
// @Param jot body types.UpdateJotPayload true "jotID and update"
// @Success 204
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 403 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/jots [patch]
func (h *Handler) handleUpdateJotByJotID(w http.ResponseWriter, req *http.Request) {
	userID := auth.GetUserIDFromContext(req.Context())
	if userID == -1 {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("couldn't find user id"))
		return
	}

	var jotToUpdate types.UpdateJotPayload
	if err := utils.ParseJSON(req, &jotToUpdate); err != nil {
		utils.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := h.jotStore.UpdateJotByJotID(jotToUpdate, int64(userID)); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
