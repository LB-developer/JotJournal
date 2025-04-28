package jots_test

import (
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lb-developer/jotjournal/service/jots"
	"github.com/lb-developer/jotjournal/utils/testutils"
)

var dbpool *pgxpool.Pool

func TestMain(m *testing.M) {
	var cleanup func()
	dbpool, cleanup = testutils.SetupDockerTest()
	defer cleanup()

	code := m.Run()
	os.Exit(code)
}

func TestGetJotsByUserID(t *testing.T) {
	store := jots.NewStore(dbpool)

	tests := []struct {
		name           string
		userID         int64
		month          int
		expectedLength int
		expectedHabits []string
	}{
		{
			name:           "ReturnsExpectedNumberOfJots",
			userID:         1,
			month:          4,
			expectedLength: 2,
		},
		{
			name:           "ReturnsEmptyWhenNoMatches",
			userID:         1,
			month:          0,
			expectedLength: 0,
		},
		{
			name:           "ReturnsOnlyJotsOfTheUser",
			userID:         1,
			month:          4,
			expectedLength: 2,
			expectedHabits: []string{"run", "walk"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jots, err := store.GetJotsByUserID(test.month, test.userID)
			if err != nil {
				t.Fatalf("Couldn't get jots, err: %s", err)
			}

			if len(jots) != test.expectedLength {
				t.Fatalf("Expected %d jots, got %d", test.expectedLength, len(jots))
			}

			// Optional habit check if expectedHabits is set
			if len(test.expectedHabits) > 0 {
				for _, habit := range test.expectedHabits {
					if _, found := jots[habit]; !found {
						t.Fatalf("Expected habit %s - not found", habit)
					}
				}
			}
		})
	}
}
