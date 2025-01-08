package user

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lb-developer/jotjournal/service/auth"
	"github.com/lb-developer/jotjournal/types"
	"github.com/lb-developer/jotjournal/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Post("/login", h.handleLogin)
	router.Post("/register", h.handleRegisterUser)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload
	err := utils.ParseJSON(r, user)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// check if the user exists
	_, err = h.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email '%s' already exists", user.Email))
		return
	}

	// hash users password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
