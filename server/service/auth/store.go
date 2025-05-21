package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lb-developer/jotjournal/config"
	uuid "github.com/satori/go.uuid"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) CreateSession(userID int64) (uuid.UUID, error) {
	u1 := uuid.NewV4()
	expiration := config.Envs.RefreshExpirationInSeconds

	query := `
		INSERT INTO 
			sessions (user_id, refresh_token, expires_at)
		VALUES
			($, $, $, CURRENT_TIMESTAMP)
	`

	tx, err := s.db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return uuid.UUID{}, err
	}
	defer tx.Rollback(context.Background())
	tag, err := tx.Exec(context.Background(), query, userID, u1, expiration)
	if err != nil {
		return uuid.UUID{}, err
	}
	if tag.RowsAffected() != 1 {
		return uuid.UUID{}, fmt.Errorf("Session was not created in DB")
	}
	if err := tx.Commit(context.Background()); err != nil {
		return uuid.UUID{}, fmt.Errorf("Transaction failed")
	}
	
	return u1, nil
}

func (s *Store) ValidateSession(userID int64, refreshToken uuid.UUID) (bool, error) {

	// does a current session exist that
	// > matches the given refreshToken && userID
	// > has not been rotated
	query := `
		SELECT
			CURRENT_TIMESTAMP < expires_at
		FROM 
			sessions
		WHERE
			user_id = $
		AND
			refresh_token = $
		AND 
			rotated = false
	`

	row := s.db.QueryRow(context.Background(), query, userID, refreshToken)

	var valid bool
	err := row.Scan(&valid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil // no matching session
		}
		return false, err // other scan error
	}	

	return valid, nil
}

