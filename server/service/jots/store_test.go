package jots_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lb-developer/jotjournal/service/jots"
	"github.com/lb-developer/jotjournal/service/user"
	"github.com/lb-developer/jotjournal/types"
	"github.com/lb-developer/jotjournal/utils/testutils"
)

func TestStore(t *testing.T) {
	dbpool, cleanup := testutils.SetupDockerTest()
	defer cleanup()

	store := jots.NewStore(dbpool)
	userStore := user.NewStore(dbpool)

	// Create a new user for testing
	testUserID := 999
	_, err := userStore.CreateUser(types.User{ID: testUserID})
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name       string
		testFunc   func(*testing.T, *jots.Store, *pgxpool.Conn)
		wantErr    bool
		wantResult any
	}{
		{
			name: "GetJotsByUserID",
			testFunc: func(t *testing.T, store *jots.Store, tx *pgxpool.Conn) {
				userID := int64(testUserID)
				month := 1
				gotJots, err := store.GetJotsByUserID(month, userID)
				if err != nil {
					t.Fatal(err)
				}
				if len(gotJots) != 0 {
					t.Errorf("expected 0 jots, got %d", len(gotJots))
				}
			},
			wantErr: false,
		},
		{
			name: "UpdateJotByJotID",
			testFunc: func(t *testing.T, store *jots.Store, tx *pgxpool.Conn) {
				jotID := 1
				jot := types.UpdateJotPayload{JotID: jotID, IsCompleted: true}
				err := store.UpdateJotByJotID(jot, 1)
				if err != nil {
					t.Fatal(err)
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := dbpool.Acquire(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			defer db.Release()

			tt.testFunc(t, store, db)
		})
	}
}
