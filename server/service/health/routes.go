package health

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lb-developer/jotjournal/utils"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Group(func(r chi.Router) {
		r.Route("/health", func(r chi.Router) {
			r.Get("/", h.handleHealthCheck)
		})
	})
}

// @Summary Health check
// @Description returns 200 when requested
// @Tags health
// @Success 200
// @Router /api/v1/health [get]
func (h *Handler) handleHealthCheck(w http.ResponseWriter, req *http.Request) {
	utils.WriteJSON(w, http.StatusOK, nil)
}
