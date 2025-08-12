package jots

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lb-developer/jotjournal/types"
	"github.com/lb-developer/jotjournal/utils"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) GetJotsByUserID(month int, year int, userID int64) (types.Jots, error) {
	query := `
	SELECT id, habit, date, is_completed 
	FROM
		jots
	WHERE
		user_id = $1
	AND 
		EXTRACT(MONTH FROM jots."date") = $2
	AND
		EXTRACT(YEAR FROM jots."date") = $3
	ORDER BY
		date
	`

	rows, err := s.db.Query(context.Background(), query, userID, month, year)
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

func (s *Store) CreateJotsForMonth(userID int64, habit string, year int, month int) ([]types.Jot, error) {
	// does habit already exist for date given
	var exists bool
	checkQuery := `
		SELECT EXISTS (
			SELECT 1 FROM jots
			WHERE user_id = $1 AND habit = $2
			AND date >= $3 AND date < $4
		)`

	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0) // first of next month
	ctx := context.Background()

	err := s.db.QueryRow(context.Background(), checkQuery, userID, habit, start, end).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("habit '%s' already exists for %d-%02d", habit, year, month)
	}

	// habit does not currently exist
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	days := utils.DaysIn(month, year)
	query := "INSERT INTO jots (user_id, habit, date) VALUES "
	args := []any{}
	argIndex := 1

	for i := 1; i <= days; i++ {
		date := time.Date(year, time.Month(month), i, 0, 0, 0, 0, time.UTC)
		query += fmt.Sprintf("($%d, $%d, $%d),", argIndex, argIndex+1, argIndex+2)
		args = append(args, userID, habit, date)
		argIndex += 3
	}
	query = query[:len(query)-1] // remove trailing comma
	query += "RETURNING id, habit, date, is_completed"

	rows, err := tx.Query(ctx, query, args...)

	var jots []types.Jot
	for rows.Next() {
		var jot types.Jot
		err := rows.Scan(&jot.ID, &jot.Habit, &jot.Date, &jot.IsCompleted)
		if err != nil {
			tx.Rollback(ctx)
			return nil, err
		}
		jots = append(jots, jot)
	}

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return nil, fmt.Errorf("commit error: %v", err)
	}

	return jots, nil
}

func (s *Store) DeleteJotsByHabit(habit string, month int, year int, userID int64) error {
	query := `
	DELETE
	FROM
		jots
	WHERE
		user_id = $1
	AND 
		habit = $2
	AND
		EXTRACT(MONTH FROM jots."date") = $3
	AND
        EXTRACT(YEAR FROM jots."date") = $4
	`

	tag, err := s.db.Exec(context.Background(), query, userID, habit, month, year)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Couldn't delete jots: %s", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("No jots to delete")
	}

	return nil
}
