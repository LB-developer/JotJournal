package jots

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
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
			r.Post("/", h.handleCreateJot)
			r.Delete("/", h.handleDeleteJotsByHabit)
		})
	})
}

// @Summary Get jots for the authenticated user
// @Description Retrieves all jots associated with the authenticated user based on their ID for the given month and year
// @Tags jots
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "JWT access token for authentication"
// @Param month query string true "month to get jots for"
// @Param year query string true "year to get jots for"
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

	// TODO: handle int extraction in util function
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

	year := req.URL.Query().Get("year")
	if year == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no year in request"))
		return
	}

	intYear, err := strconv.Atoi(year)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	jots, err := h.jotStore.GetJotsByUserID(intMonth, intYear, int64(userID))
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

// @Summary Creates a new jot for the authenticated user
// @Description Creates a jot with name and date for the authenticated user
// @Tags jots
// @Security BearerAuth
// @Param Authorization header string true "JWT access token for authentication"
// @Param payload body types.CreateJotPayload true "name and date"
// @Success 201 {object} []types.Jot
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 403 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/jots [post]
func (h *Handler) handleCreateJot(w http.ResponseWriter, req *http.Request) {
	userID := auth.GetUserIDFromContext(req.Context())
	if userID == -1 {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("couldn't find user id"))
		return
	}

	var payload types.CreateJotPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	jots, err := h.jotStore.CreateJotsForMonth(int64(userID), payload.Name, payload.Year, payload.Month)
	if err != nil || jots == nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, jots)
}

// @Summary Delete jots for a partificular month/year by habit for the authenticated user
// @Description Deletes all jots associated with the authenticated user based on the habit for the given month and year
// @Tags jots
// @Security BearerAuth
// @Param Authorization header string true "JWT access token for authentication"
// @Param habit query string true "habit to delete"
// @Param month query string true "month of specified habit to delete"
// @Param year query string true "year of specified habit to delete"
// @Success 200
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 403 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/jots [delete]
func (h *Handler) handleDeleteJotsByHabit(w http.ResponseWriter, req *http.Request) {
	userID := auth.GetUserIDFromContext(req.Context())
	if userID == -1 {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("couldn't find user id"))
		return
	}

	// set the payload
	var payload types.DeleteJotPayload
	habit := req.Header.Get("habit")
	strMonth := req.Header.Get("month")
	strYear := req.Header.Get("year")

	intMonth, err := strconv.Atoi(strMonth)
	intYear, err := strconv.Atoi(strYear)

	payload.Habit = habit
	payload.Month = intMonth
	payload.Year = intYear

	// validate the delete jot payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid payload %v\n", errors))
		return
	}

	err = h.jotStore.DeleteJotsByHabit(payload.Habit, payload.Month, payload.Year, int64(userID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("couldn't delete jots: %s", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
