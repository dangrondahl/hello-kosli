package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/dangrondahl/hello-kosli/internal/logging"
)

type helloReq struct {
	Name string `json:"name"`
}

// Hello handles POST /hello and replies "Hello, <param>" as plain text.
// Accepts JSON body {"name": "..."} or query string ?name=... (body takes precedence).
var log = logging.Nop()

// SetLogger sets the package logger for handlers. Call from main to provide
// a real logger (e.g. logging.NewStdLogger(os.Stdout)).
func SetLogger(l logging.Logger) {
	if l == nil {
		log = logging.Nop()
		return
	}
	log = l
}

type respWriter struct {
	http.ResponseWriter
	status int
}

func (rw *respWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	rw := &respWriter{ResponseWriter: w, status: http.StatusOK}
	defer func() {
		// Emit a structured log for the request
		log.Info("handled request", map[string]interface{}{
			"method":     r.Method,
			"path":       r.URL.Path,
			"remote_addr": r.RemoteAddr,
			"user_agent": r.UserAgent(),
			"status":     rw.status,
		})
	}()
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
		http.Error(rw, "missing name", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write([]byte("Hello, " + name))
}