package session

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lb-developer/jotjournal/config"
	"github.com/valkey-io/valkey-glide/go/api"
)

type Store struct {
	db *pgxpool.Pool
	cache api.GlideClientCommands
}

func NewStore(db *pgxpool.Pool, cache api.GlideClientCommands) *Store {
	return &Store{db: db, cache: cache}
}

func (s *Store) CacheSessionToken(sessionToken string, sessionID string) (string, error) {
	set, err := s.cache.Set(sessionToken, sessionID)
	return set, err
}

func (s *Store) ValidateSessionToken(refreshToken string) (string, error) {
	res, err := s.cache.Get(refreshToken)
	if res.IsNil() {
		return "", fmt.Errorf("No current session found...")
	}

	if err != nil {
		return "", err
	}

	// user has a current session
	return res.Value(), nil
}


func (s *Store) CreateSession(userID int64) (string, error) {
	expiration := time.Now().Add(time.Second * time.Duration(config.Envs.RefreshExpirationInSeconds))
	
	query := `
		INSERT INTO 
			sessions (user_id, expires_at, created_at)
		VALUES
			($1, $2, CURRENT_TIMESTAMP)
		RETURNING
			id
	`

	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	var sessionID string 

	err = tx.QueryRow(ctx, query, userID, expiration).Scan(&sessionID)
	if err != nil {
		return "", fmt.Errorf("QueryRow failed: %s", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("Transaction failed, err: %s", err)
	}
	
	return sessionID, nil
}

func (s *Store) ValidateSession(userID int64, refreshToken string) (bool, error) {

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


func (s *Store) DestroySession(userID int64, sessionToken string) (bool, error) {
	query := `
		DELETE FROM
			sessions
		WHERE
			user_id = $1
		AND 
			id = $2
	`

	// get sessionID from the cache
	sessionID, err := s.cache.Get(sessionToken)
	if sessionID.IsNil() {
		return false, fmt.Errorf("Session was not found in the cache")
	}

	if err != nil {
		return false, err
	}

	// remove session from cache
	s.cache.Del([]string{sessionToken})

	// remove session from db
	tag, err := s.db.Exec(context.Background(), query, userID, sessionID)


	if tag.RowsAffected() != 1 {
		return false, fmt.Errorf("Session for user %v was not deleted because, err: %s", userID, err)
	}

	if err != nil {
		return false, err
	}

	// session successfully destroyed
	return true, nil
}
