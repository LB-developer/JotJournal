package types

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ErrorResponse struct {
	Error string `json:"error" example:"something went wrong"`
}

type SuccessfulLoginResponse struct {
	SessionToken string `json:"sessionToken" example:"header.payload.signature"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=100"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user User) error
}

type Task struct {
	ID          int       `json:"id" validate:"required" faker:"oneof: 1, 2"`
	Monthly     bool      `json:"monthly" faker:"-"`
	Weekly      bool      `json:"weekly" faker:"-"`
	Daily       bool      `json:"daily" faker:"-"`
	Deadline    time.Time `json:"deadline" validate:"required" example:"2006-01-02T15:04:00Z" faker:"-"`
	Description string    `json:"description" validate:"required" faker:"sentence"`
	IsCompleted bool      `json:"isCompleted" faker:"-"`
	UserID      int       `json:"userID" validate:"required" faker:"oneof: 1, 2"`
}

type NewTask struct {
	Monthly     bool      `json:"monthly" faker:"-"`
	Weekly      bool      `json:"weekly"  faker:"-"`
	Daily       bool      `json:"daily" faker:"-"`
	Deadline    time.Time `json:"deadline" validate:"required" example:"2025-05-01T00:00:00Z" faker:"timestamp"`
	Description string    `json:"description" validate:"required" faker:"sentence"`
}

type (
	Jots map[string][]Jot
	Jot  struct {
		ID          int       `json:"id" validate:"required" faker:"-"`
		Habit       string    `json:"habit" validate:"required" example:"workout" faker:"sentence"`
		Date        time.Time `json:"date" validate:"required" example:"2006-01-02T15:04:00Z" faker:"-"`
		IsCompleted bool      `json:"isCompleted" faker:"-"`
	}
	UpdateJotPayload struct {
		JotID       string `json:"jotID"`
		IsCompleted bool   `json:"isCompleted"`
	}
)

type TaskIDToDelete struct {
	ID int
}

type TaskStore interface {
	GetTasksByUserID(userID int64) ([]Task, error)
	UpdateTaskByTaskID(editedTask Task) (Task, error)
	DeleteTaskByTaskID(taskId TaskIDToDelete, userID int64) error
	CreateTask(task NewTask, userID int64) (int, error)
}

type JotStore interface {
	GetJotsByUserID(month int, userID int64) (Jots, error)
	UpdateJotByJotID(jot UpdateJotPayload, userID int64) error
}

type SessionStore interface {
	CreateSession(userID int64) (string, error)
	ValidateSession(userID int64, uuid string) (bool, error)
	CacheSessionToken(sessionToken string, sessionID string) (string, error)
	ValidateSessionToken(sessionToken string) (string, error)
}

type RefreshTokenPayload struct {
	RefreshToken uuid.UUID `json:"refreshToken" example:"abc-123-xyz-123"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"accessToken" example:"header.payload.signature"`
}
