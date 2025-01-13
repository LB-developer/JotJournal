package tasks

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lb-developer/jotjournal/types"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) GetTasksByUserID(user_id int64) ([]types.Task, error) {
	var tasks []types.Task
	return tasks, nil
}
func (s *Store) UpdateTaskByTaskID(taskId int64) (types.Task, error) {
	return types.Task{}, nil
}

func (s *Store) DeleteTaskByTaskID(taskId int64) error {
	return nil
}

func (s *Store) CreateTask(userId int64) error {
	return nil
}
