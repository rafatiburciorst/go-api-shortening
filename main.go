package main

import (
	"encurtador/api"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code: ")
		return
	}
	slog.Info("All systems offline")
}

func run() error {
	db := make(map[string]string)
	handler := api.NewHandler(db)
	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         "localhost:8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	slog.Info("HTTP Server is running")
	return nil
}
