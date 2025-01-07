package user

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/lb-developer/jotjournal/types"
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
}
