package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type helloReq struct {
	Name string `json:"name"`
}

// Hello handles POST /hello and replies "Hello, <param>" as plain text.
// Accepts JSON body {"name": "..."} or query string ?name=... (body takes precedence).
func Hello(w http.ResponseWriter, r *http.Request) {
	name := ""
	// Try JSON body
	if r.Body != nil {
		defer r.Body.Close()
		b, _ := io.ReadAll(io.LimitReader(r.Body, 1<<20)) // 1MB cap
		if len(b) > 0 {
			var req helloReq
			if err := json.Unmarshal(b, &req); err == nil {
				name = strings.TrimSpace(req.Name)
			}
		}
	}
	// Fallback to query
	if name == "" {
		name = strings.TrimSpace(r.URL.Query().Get("name"))
	}
	if name == "" {
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, " + name))
}

func parseNameFromBody(body io.Reader) (string, error) {
	var req helloReq
	if err := json.NewDecoder(body).Decode(&req); err != nil {
		return "", err
	}
	n := strings.TrimSpace(req.Name)
	if n == "" {
		return "", errors.New("empty")
	}
	return n, nil
}
