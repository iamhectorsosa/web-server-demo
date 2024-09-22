package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iamhectorsosa/web-server-demo/internal/memorystore"
	"github.com/iamhectorsosa/web-server-demo/internal/store"
)

func TestServer(t *testing.T) {
	cleanup := EnvLoad()
	defer cleanup()
	initialUsers := []store.User{
		{Id: "1", Email: "sosa@webscope.io"},
	}
	store := memorystore.New(initialUsers...)
	server := New(store)

	t.Run("returns the server's health", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/api/health", nil)
		if err != nil {
			t.Errorf("error creating new request: %v", err)
		}
		response := httptest.NewRecorder()
		server.Handler.ServeHTTP(response, request)

		AssertStatus(t, http.StatusOK, response.Code)
		AssertContentType(t, "application/json", response.Header())
		AssertResponseBody(t, "{\"status\":\"OK\"}", response.Body.String())
	})

	t.Run("returns the store's users", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/api/users", nil)
		if err != nil {
			t.Errorf("error creating new request: %v", err)
		}
		response := httptest.NewRecorder()
		server.Handler.ServeHTTP(response, request)

		AssertStatus(t, http.StatusOK, response.Code)
		AssertContentType(t, "application/json", response.Header())
		AssertResponseBody(t, "[{\"id\":\"1\",\"email\":\"sosa@webscope.io\"}]", response.Body.String())
	})

	t.Run("returns a store user", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/api/users/1", nil)
		if err != nil {
			t.Errorf("error creating new request: %v", err)
		}
		response := httptest.NewRecorder()
		server.Handler.ServeHTTP(response, request)

		AssertStatus(t, http.StatusOK, response.Code)
		AssertContentType(t, "application/json", response.Header())
		AssertResponseBody(t, "{\"id\":\"1\",\"email\":\"sosa@webscope.io\"}", response.Body.String())
	})
}
