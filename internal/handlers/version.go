package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dangrondahl/hello-kosli/internal/version"
)

type versionResp struct {
	GitSHA string `json:"git_sha"`
}

func Version(w http.ResponseWriter, r *http.Request) {
	// Emit structured log for the request. `log` is the package logger set by SetLogger.
	log.Info("handled request", map[string]interface{}{
		"method": r.Method,
		"path":   r.URL.Path,
		"remote": r.RemoteAddr,
	})

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(versionResp{GitSHA: version.GitSHA})
}
