package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер, который пишет в стандартный вывод
	logger := log.New(os.Stdout, "app: ", log.LstdFlags|log.Lshortfile)

	// Создаем сервер, передавая логгер
	srv := server.NewServer(logger)

	// Запускаем сервер
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("Ошибка запуска сервера: ", err)
	}
}
