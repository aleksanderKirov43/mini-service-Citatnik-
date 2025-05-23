package app

import (
	"log"
	"net/http"

	"mini-service-Citatnik-/internal/router"
)

func Run() error {
	r := router.NewRouter()

	log.Println("Сервер запущен на порту :8080")
	return http.ListenAndServe(":8080", r)
}
