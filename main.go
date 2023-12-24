package main

import (
	"log"
	"net/http"
	"os"
	"profiley/handlers"
	"time"
)

func main() {
	l := log.New(os.Stdout, "profile-api", log.LstdFlags)
	hh := handlers.NewData(l)

	serve_mux := http.NewServeMux()
	serve_mux.Handle("/", hh)

	server := &http.Server{
		Addr:         ":3000",
		Handler:      serve_mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	server.ListenAndServe()

}
