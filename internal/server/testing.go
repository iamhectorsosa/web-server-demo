package server

import (
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func EnvLoad() (cleanup func()) {
	envDir := ".env"
	_ = godotenv.Write(map[string]string{"PORT": "8080"}, envDir)

	return func() {
		os.RemoveAll(envDir)
	}
}

func AssertResponseBody(t testing.TB, want, got string) {
	t.Helper()

	if want != got {
		t.Errorf("response body error: want %s, got %s", want, got)
	}
}

func AssertStatus(t testing.TB, want, got int) {
	t.Helper()

	if want != got {
		t.Errorf("response code error: want %d, got %d", want, got)
	}
}

func AssertContentType(t testing.TB, want string, header http.Header) {
	t.Helper()
	if want != header.Get("Content-Type") {
		t.Errorf("content type error: want %s, got %v", want, header)
	}
}
