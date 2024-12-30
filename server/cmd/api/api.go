package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lb-developer/jotjournal/handlers/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: address,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()

	subrouter := chi.NewRouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	r.Mount("/api/v1", subrouter)

	log.Printf("Listening on %s", s.addr)
	return http.ListenAndServe(s.addr, r)
}
