package main

import (
	"book-shelf/config"
	"errors"
	"io"
	"log"
	"net/http"
)

func main() {
	c := config.NewConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  c.Server.ReadTimeout,
		WriteTimeout: c.Server.WriteTimeout,
		IdleTimeout:  c.Server.IdleTimeout,
	}

	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server startup failed")
	}
}

func hello(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "Hello, World!")
	if err != nil {
		return
	}
}
