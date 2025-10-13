package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dangrondahl/hello-kosli/internal/version"
)

func TestVersion(t *testing.T) {
	version.GitSHA = "abc123" // override for test

	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	rec := httptest.NewRecorder()

	Version(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusOK)
	}
	var out struct {
		GitSHA string `json:"git_sha"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out.GitSHA != "abc123" {
		t.Fatalf("git_sha: got %q, want %q", out.GitSHA, "abc123")
	}
}
