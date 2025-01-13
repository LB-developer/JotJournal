package tasks

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lb-developer/jotjournal/types"
	"github.com/lb-developer/jotjournal/utils"
)

type Handler struct {
	store types.TaskStore
}

func NewHandler(store types.TaskStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
}

func (h *Handler) handleGetTasksByUserId(w http.ResponseWriter, req *http.Request) {
}

func (h *Handler) handleCreateTask(w http.ResponseWriter, req *http.Request) {
}

func (h *Handler) handleUpdateTaskByTaskId(w http.ResponseWriter, req *http.Request) {
}

func (h *Handler) handleDeleteTaskByTaskId(w http.ResponseWriter, req *http.Request) {
}
