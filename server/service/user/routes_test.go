package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/lb-developer/jotjournal/types"
)

var payloadAndStatus = map[string]struct {
	input          types.RegisterUserPayload
	expectedStatus int
}{
	"no email": {
		input: types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "",
			Password:  "123"},
		expectedStatus: http.StatusUnprocessableEntity,
	},
	"invalid email": {
		input: types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "abc",
			Password:  "123"},
		expectedStatus: http.StatusUnprocessableEntity,
	},
	"valid email": {
		input: types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "valid@gmail.com",
			Password:  "123"},
		expectedStatus: http.StatusOK,
	},
}

func TestUserServiceHandlers(t *testing.T) {
	mockUserStore := &mockUserStore{}
	mockSessionStore := &mockSessionStore{}
	handler := NewHandler(mockUserStore, mockSessionStore)

	for name, test := range payloadAndStatus {
		t.Run(name, func(t *testing.T) {
			payload := test.input
			marshalled, _ := json.Marshal(payload)

			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
			rr := httptest.NewRecorder()

			router := chi.NewRouter()
			router.Post("/register", handler.handleRegisterUser)

			router.ServeHTTP(rr, req)

			if rr.Code != test.expectedStatus {
				t.Errorf("got %d wanted %d", rr.Code, test.expectedStatus)
			}
		})
	}
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) (int, error) {
	return 0, nil
}

type mockSessionStore struct{}

func (m *mockSessionStore) CreateSession(userID int64) (string, error) {
	return "", nil
}

func (m *mockSessionStore) ValidateSession(userID int64, refreshToken string) (bool, error) {
	return true, nil
}

func (s *mockSessionStore) CacheSessionToken(sessionToken string, sessionID string) (string, error) {
	return "", nil
}

func (m *mockSessionStore) ValidateSessionToken(sessionToken string) (string, error) {
	return "", nil
}

func (m *mockSessionStore) DestroySession(userID int64, sessionToken string) (bool, error) {
	return true, nil
}

func (m *mockSessionStore) ClearSessionFromCache(sessionToken string) (bool, error) {
	return true, nil
}
