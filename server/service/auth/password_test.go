package auth

import "testing"

var tests = map[string]struct {
	passwordOne string
	passwordTwo []byte
	expected    bool
}{
	"same_password": {
		passwordOne: "samepassword",
		passwordTwo: []byte("samepassword"),
		expected:    true,
	},
	"different_password": {
		passwordOne: "password",
		passwordTwo: []byte("differentPassword"),
		expected:    false,
	},
}

func TestHashPassword(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			hashed, _ := HashPassword(test.passwordOne)

			actual := ComparePasswords(hashed, test.passwordTwo)

			if test.expected != actual {
				t.Errorf("passwords do not match")
			}
		})
	}
}
