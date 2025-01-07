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
	user := new(types.User)
	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (s *Store) CreateUser(user types.User) error {
	return nil
}
