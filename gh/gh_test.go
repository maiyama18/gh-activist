package gh

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name                    string
		user                    string
		token                   string
		repo                    string
		file                    string
		expectedAuthHeaderValue string
	}{
		{
			name:                    "success",
			user:                    "USER",
			token:                   "TOKEN",
			repo:                    "REPO",
			file:                    "FILE",
			expectedAuthHeaderValue: "token TOKEN",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewClient(test.user, test.token, test.repo, test.file)
			assert.Equal(t, test.expectedAuthHeaderValue, c.AuthHeaderValue)
		})
	}
}

func TestClient_Commit(t *testing.T) {
	tests := []struct {
		name           string
		expectedErrMsg string
	}{
		{
			name: "success",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

		})
	}
}
