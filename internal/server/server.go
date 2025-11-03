package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	*log.Logger
	*http.Server
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	// Регистрация маршрутов с обработчиками
	mux.HandleFunc("/", handlers.FirstHandler)
	mux.HandleFunc("/upload", handlers.SecondHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: server,
	}
}
