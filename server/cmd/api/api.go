package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lb-developer/jotjournal/docs"
	"github.com/lb-developer/jotjournal/service/health"
	"github.com/lb-developer/jotjournal/service/jots"
	"github.com/lb-developer/jotjournal/service/session"
	"github.com/lb-developer/jotjournal/service/tasks"
	"github.com/lb-developer/jotjournal/service/user"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/valkey-io/valkey-glide/go/api"
)

// @title JotJournal API

// @version 1.0
// @description This is the RESTful API backend for the JotJournal app. It handles task management & user data.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
type APIServer struct {
	addr  string
	db    *pgxpool.Pool
	cache api.GlideClientCommands
}

func NewAPIServer(address string, db *pgxpool.Pool, cache api.GlideClientCommands) *APIServer {
	return &APIServer{
		addr:  address,
		db:    db,
		cache: cache,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	subrouter := chi.NewRouter()

	healthHandler := health.NewHandler()
	healthHandler.RegisterRoutes(subrouter)

	sessionStore := session.NewStore(s.db, s.cache)
	userStore := user.NewStore(s.db)
	taskStore := tasks.NewStore(s.db)
	jotStore := jots.NewStore(s.db)

	userHandler := user.NewHandler(userStore, sessionStore)
	taskHandler := tasks.NewHandler(taskStore, userStore, sessionStore)
	jotHandler := jots.NewHandler(jotStore, userStore, sessionStore)

	jotHandler.RegisterRoutes(subrouter)
	taskHandler.RegisterRoutes(subrouter)
	userHandler.RegisterRoutes(subrouter)

	r.Mount("/api/v1", subrouter)

	// Register the Swagger handler
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	log.Printf("Listening on %s", s.addr)
	return http.ListenAndServe(s.addr, r)
}
