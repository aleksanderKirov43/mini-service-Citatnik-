package app

import (
	"log"
	"mini-service-Citatnik-/internal/handler"
	"net/http"

	"mini-service-Citatnik-/internal/router"
	"mini-service-Citatnik-/internal/service"
)

func Run() error {
	s := service.NewQuoteService()
	h := handler.NewQuoteHandler(s)

	r := router.NewRouter(h)

	log.Println("Сервер запущен на порту :8080")
	return http.ListenAndServe(":8080", r)
}
