package types

import "time"

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
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
	ID          int       `json:"id" faker:"oneof: 1, 2"`
	Monthly     bool      `json:"monthly" faker:"-"`
	Weekly      bool      `json:"weekly" faker:"-"`
	Daily       bool      `json:"daily" faker:"-"`
	Deadline    time.Time `json:"deadline" faker:"timestamp"`
	Description string    `json:"description" faker:"sentence"`
	IsCompleted bool      `json:"isCompleted" faker:"-"`
	UserID      int       `json:"userId" faker:"oneof: 1, 2"`
}

type NewTask struct {
	Monthly     bool      `json:"monthly" validate:"required" faker:"-"`
	Weekly      bool      `json:"weekly"  validate:"required" faker:"-"`
	Daily       bool      `json:"daily" validate:"required" faker:"-"`
	Deadline    time.Time `json:"deadline" validate:"required" faker:"timestamp"`
	Description string    `json:"description" validate:"required" faker:"sentence"`
}

type TaskStore interface {
	GetTasksByUserID(userId int64) ([]Task, error)
	UpdateTaskByTaskID(taskId int64) (Task, error)
	DeleteTaskByTaskID(taskId int64) error
	CreateTask(task NewTask, userId int64) (int, error)
}
