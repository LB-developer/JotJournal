package types

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type RefreshTokenPayload struct {
	RefreshToken uuid.UUID `json:"refreshToken" example:"abc-123-xyz-123"`
}

type SessionTokenResponse struct {
	SessionToken string `json:"sessionToken" example:"header.payload.signature"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"something went wrong"`
}

type SuccessfulLoginResponse struct {
	SessionToken string       `json:"sessionToken" example:"header.payload.signature"`
	User         UserResponse `json:"user"`
}

type SessionStore interface {
	CreateSession(userID int64) (string, error)
	ValidateSession(userID int64, uuid string) (bool, error)
	CacheSessionToken(sessionToken string, sessionID string) (string, error)
	ValidateSessionToken(sessionToken string) (string, error)
	DestroySession(userID int64, sessionToken string) (bool, error)
	ClearSessionFromCache(sessionToken string) (bool, error)
}

type (
	UserStore interface {
		GetUserByEmail(email string) (*User, error)
		GetUserByID(id int) (*User, error)
		CreateUser(user User) (int, error)
	}

	User struct {
		ID        int       `json:"id"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Email     string    `json:"email"`
		Password  string    `json:"-"`
		CreatedAt time.Time `json:"createdAt"`
	}

	UserResponse struct {
		ID        int    `json:"ID"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	}

	RegisterUserPayload struct {
		FirstName string `json:"firstName" validate:"required"`
		LastName  string `json:"lastName"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=3,max=100"`
	}

	LoginUserPayload struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LogoutUserPayload struct {
		SessionToken string `json:"sessionToken" example:"header.payload.signature"`
	}
)

type (
	TaskStore interface {
		GetTasksByUserID(userID int64) ([]Task, error)
		UpdateTaskByTaskID(editedTask Task) (Task, error)
		DeleteTaskByTaskID(taskId TaskIDToDelete, userID int64) error
		CreateTask(task NewTask, userID int64) (int, error)
	}

	Task struct {
		ID          int       `json:"id" validate:"required" faker:"oneof: 1, 2"`
		Monthly     bool      `json:"monthly" faker:"-"`
		Weekly      bool      `json:"weekly" faker:"-"`
		Daily       bool      `json:"daily" faker:"-"`
		Deadline    time.Time `json:"deadline" validate:"required" example:"2006-01-02T15:04:00Z" faker:"-"`
		Description string    `json:"description" validate:"required" faker:"sentence"`
		IsCompleted bool      `json:"isCompleted" faker:"-"`
		UserID      int64     `json:"userID" validate:"required" faker:"oneof: 1, 2"`
	}

	TaskIDToDelete struct {
		ID int
	}

	NewTask struct {
		Monthly     bool      `json:"monthly" faker:"-"`
		Weekly      bool      `json:"weekly"  faker:"-"`
		Daily       bool      `json:"daily" faker:"-"`
		Deadline    time.Time `json:"deadline" validate:"required" example:"2025-05-01T00:00:00Z" faker:"timestamp"`
		Description string    `json:"description" validate:"required" faker:"sentence"`
	}
)

type (
	JotStore interface {
		GetJotsByUserID(month int, year int, userID int64) (Jots, error)
		UpdateJotByJotID(jot UpdateJotPayload, userID int64) error
		CreateJotsForMonth(userID int64, habit string, year int, month int) ([]Jot, error)
		DeleteJotsByHabit(habit string, month int, year int, userID int64) error
	}

	Jots map[string][]Jot
	Jot  struct {
		ID          int       `json:"id" validate:"required" faker:"-"`
		Habit       string    `json:"habit" validate:"required" example:"workout" faker:"sentence"`
		Date        time.Time `json:"date" validate:"required" example:"2006-01-02T15:04:00Z" faker:"-"`
		IsCompleted bool      `json:"isCompleted" faker:"-"`
	}
	UpdateJotPayload struct {
		JotID       int  `json:"jotID"`
		IsCompleted bool `json:"isCompleted"`
	}
	CreateJotPayload struct {
		Name  string `json:"name" validate:"required" example:"Run"`
		Month int    `json:"month" validate:"required" example:"04"`
		Year  int    `json:"year" validate:"required" example:"2025"`
	}
	DeleteJotPayload struct {
		Habit string `json:"habit" validate:"required" example:"workout"`
		Month int    `json:"month" validate:"required" example:"04"`
		Year  int    `json:"year" validate:"required" example:"2025"`
	}
)
