package user

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lb-developer/jotjournal/types"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	query := `
	SELECT
			*
	FROM
			users
	WHERE
			email = $1
	`
	rows, err := s.db.Query(context.Background(), query, email)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query for user failed during scan: %v\n", err)
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *Store) GetUserByID(userId int) (*types.User, error) {
	query := `
	SELECT
			*
	FROM
			users
	WHERE
			id = $1
	`
	rows, err := s.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query for user failed during scan: %v\n", err)
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *Store) CreateUser(user types.User) error {
	query := `
	INSERT INTO
		users
		(first_name, last_name, email, password)
	VALUES
		($1, $2, $3, $4)
	`

	_, err := s.db.Exec(context.Background(), query, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("failed to acquire connection, err: \n%v", err)
	}

	return nil
}

func scanRowIntoUser(rows pgx.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

