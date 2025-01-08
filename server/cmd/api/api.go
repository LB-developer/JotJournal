package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lb-developer/jotjournal/service/user"
)

type APIServer struct {
	addr string
	db   *pgxpool.Pool
}

func NewAPIServer(address string, db *pgxpool.Pool) *APIServer {
	return &APIServer{
		addr: address,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()

	subrouter := chi.NewRouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	r.Mount("/api/v1", subrouter)

	log.Printf("Listening on %s", s.addr)
	return http.ListenAndServe(s.addr, r)
}
