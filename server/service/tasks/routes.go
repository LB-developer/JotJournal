package tasks

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lb-developer/jotjournal/service/auth"
	"github.com/lb-developer/jotjournal/types"
	"github.com/lb-developer/jotjournal/utils"
)

type Handler struct {
	taskStore types.TaskStore
	userStore types.UserStore
}

func NewHandler(taskStore types.TaskStore, userStore types.UserStore) *Handler {
	return &Handler{taskStore: taskStore, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Group(func(r chi.Router) {
		r.Use(auth.ProtectedRoute(h.userStore))

		r.Route("/tasks", func(r chi.Router) {
			r.Get("/", h.handleGetTasksByUserId)
		})
	})
}

func (h *Handler) handleGetTasksByUserId(w http.ResponseWriter, req *http.Request) {
}

func (h *Handler) handleCreateTask(w http.ResponseWriter, req *http.Request) {
}

func (h *Handler) handleUpdateTaskByTaskId(w http.ResponseWriter, req *http.Request) {
}

func (h *Handler) handleDeleteTaskByTaskId(w http.ResponseWriter, req *http.Request) {
}
