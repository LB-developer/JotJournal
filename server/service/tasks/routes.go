package tasks

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
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
			r.Put("/", h.handleCreateTask)
			r.Patch("/", h.handleUpdateTask)
			r.Delete("/", h.handleDeleteTaskByTaskId)
		})
	})
}

func (h *Handler) handleGetTasksByUserId(w http.ResponseWriter, req *http.Request) {
	userID := auth.GetUserIDFromContext(req.Context())
	tasks, err := h.taskStore.GetTasksByUserID(int64(userID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, tasks)
}

func (h *Handler) handleCreateTask(w http.ResponseWriter, req *http.Request) {
	var newTask types.NewTask
	err := utils.ParseJSON(req, &newTask)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the new task payload
	if err := utils.Validate.Struct(newTask); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid payload %v\n", errors))
		return
	}

	userID := auth.GetUserIDFromContext(req.Context())

	taskId, err := h.taskStore.CreateTask(newTask, int64(userID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, taskId)
}

func (h *Handler) handleUpdateTask(w http.ResponseWriter, req *http.Request) {
	userId := auth.GetUserIDFromContext(req.Context())

	var editedTask types.Task
	if err := utils.ParseJSON(req, &editedTask); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	editedTask.UserID = userId

	// validate the edited task payload
	if err := utils.Validate.Struct(editedTask); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid payload %v\n", errors))
		return
	}

	updatedTask, err := h.taskStore.UpdateTaskByTaskID(editedTask)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, updatedTask)
}

func (h *Handler) handleDeleteTaskByTaskId(w http.ResponseWriter, req *http.Request) {
	userId := auth.GetUserIDFromContext(req.Context())

	var taskId types.TaskIDToDelete
	if err := utils.ParseJSON(req, &taskId); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.taskStore.DeleteTaskByTaskID(taskId, int64(userId))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, taskId)
}
