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
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(versionResp{GitSHA: version.GitSHA})
}
