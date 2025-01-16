package tasks

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lb-developer/jotjournal/types"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) GetTasksByUserID(userId int64) ([]types.Task, error) {
	query := `
	SELECT 
		* 
	FROM 
		tasks
	WHERE
		$1 = user_id
	`

	rows, err := s.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}

	var tasks []types.Task
	for rows.Next() {
		var task types.Task
		if err := rows.Scan(
			&task.ID,
			&task.Monthly,
			&task.Weekly,
			&task.Daily,
			&task.Deadline,
			&task.Description,
			&task.IsCompleted,
			&task.UserID,
		); err != nil {
			return nil, fmt.Errorf("Couldn't scan into tasks, error: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *Store) UpdateTaskByTaskID(taskId int64) (types.Task, error) {
	return types.Task{}, nil
}

func (s *Store) DeleteTaskByTaskID(taskId int64) error {
	query := `
	DELETE FROM 
		tasks
	WHERE
		id = $1
	`

	deleted, err := s.db.Exec(context.Background(), query, taskId)
	if err != nil {
		return err
	}

	if isDeleted := pgconn.CommandTag.Delete(deleted); !isDeleted {
		return fmt.Errorf("Couldn't confirm task was deleted from db")
	}

	return nil
}

func (s *Store) CreateTask(task types.NewTask, userId int64) (int, error) {
	query := `
	INSERT INTO tasks (monthly, weekly, daily, deadline, description, is_completed, user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id
	`

	var lastInsertId int
	err := s.db.QueryRow(
		context.Background(),
		query,
		task.Monthly,
		task.Weekly,
		task.Daily,
		task.Deadline,
		task.Description,
		false, // default is_completed
		userId,
	).Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}

	return lastInsertId, nil
}
