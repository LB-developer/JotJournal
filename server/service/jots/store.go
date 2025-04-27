package jots

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

func (s *Store) GetJotsByUserID(month int, userID int64) (types.Jots, error) {
	query := `
	SELECT id, habit, date, is_completed 
	FROM
		jots
	WHERE
		user_id = $1
	AND 
		EXTRACT(MONTH FROM jots."date") = $2
	ORDER BY
		date
	`

	rows, err := s.db.Query(context.Background(), query, userID, month)
	if err != nil {
		return types.Jots{}, err
	}

	jots := make(types.Jots)
	for rows.Next() {
		var jot types.Jot
		if err := rows.Scan(
			&jot.ID,
			&jot.Habit,
			&jot.Date,
			&jot.IsCompleted,
		); err != nil {
			return types.Jots{}, fmt.Errorf("Couldn't scan into jots, error: %v", err)
		}

		jots[jot.Habit] = append(jots[jot.Habit], jot)
	}

	return jots, nil
}

func (s *Store) UpdateJotByJotID(jot types.UpdateJotPayload, userID int64) error {
	query := `
	UPDATE 
		jots
	SET 
		is_completed = $1
	WHERE 
		id = $2
	AND
		user_id = $3
	AND
		date <= CURRENT_TIMESTAMP
	`

	tag, err := s.db.Exec(context.Background(), query, jot.IsCompleted, jot.JotID, userID)
	if err != nil {
		return fmt.Errorf("Couldn't update jot: %s", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("Cannot update jots in the future")
	}

	return nil
}
