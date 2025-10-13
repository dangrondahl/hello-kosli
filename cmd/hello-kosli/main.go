package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dangrondahl/hello-kosli/internal/handlers"
	"github.com/dangrondahl/hello-kosli/internal/logging"
)

func main() {
	// Wire the standard JSON logger to stdout
	handlers.SetLogger(logging.NewStdLogger(os.Stdout))

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handlers.Hello)
	mux.HandleFunc("/version", handlers.Version)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server: %v", err)
	}
}
