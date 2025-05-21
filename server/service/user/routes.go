package user

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/lb-developer/jotjournal/config"
	"github.com/lb-developer/jotjournal/service/auth"
	"github.com/lb-developer/jotjournal/types"
	"github.com/lb-developer/jotjournal/utils"
)

type Handler struct {
	store types.UserStore
	sessionStore types.SessionStore
}

func NewHandler(store types.UserStore, sessionStore types.SessionStore) *Handler {
	return &Handler{store: store, sessionStore: sessionStore}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Post("/login", h.handleLogin)
	router.Post("/register", h.handleRegisterUser)
}

// @Summary Logs a user in and authenticates them with a JWT access token
// @Description Authenticates a user from an email and password and begins a session
// @Tags User
// @Accepts json
// @Produce json
// @Param Login body types.LoginUserPayload true "Login input"
// @Success 200 {object} types.SuccessfulLoginResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 422 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/login [post]
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the user payload
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid payload %v\n", errors))
		return
	}

	// check if the user exists
	u, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email: '%s' or password", user.Email))
		return
	}

	// does given pw match stored pw
	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)

	// generate session token
	sessionToken, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("couldn't create access token"))
		return
	}

	// create session in database
	sessionID, err := h.sessionStore.CreateSession(int64(u.ID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("couldn't create session, err: %s", err))
		return
	}

	// add session to cache
	_, err = h.sessionStore.CacheSessionToken(sessionToken, sessionID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("couldn't cache session: %s", err))
		return
	}

	loginSuccessData := types.SuccessfulLoginResponse{
		SessionToken: sessionToken,
	}

	// successfully logged in and given token
	utils.WriteJSON(w, http.StatusOK, loginSuccessData)
}

// @Summary Registers a user in the database
// @Description Registers a user from an email and password
// @Tags User
// @Accepts json
// @Produce json
// @Param Register body types.RegisterUserPayload true "User registration input"
// @Success 200
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 422 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/register [post]
func (h *Handler) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	// unmarshal user registration payload
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the user payload
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid payload %v\n", errors))
		return
	}

	// check if the user exists
	if _, err := h.store.GetUserByEmail(user.Email); err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email '%s' already exists", user.Email))
		return
	}

	// hash users password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// create new user
	err = h.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}


	// successfully created user
	utils.WriteJSON(w, http.StatusCreated, nil)
}

