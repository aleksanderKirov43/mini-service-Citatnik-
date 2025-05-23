package app

import (
	"log"
	"net/http"

	"mini-service-Citatnik-/internal/service"
	"mini-service-Citatnik-/internal/transport/handler"
	"mini-service-Citatnik-/internal/transport/router"
)

func Run() error {
	s := service.NewQuoteService()
	h := handler.NewQuoteHandler(s)

	r := router.NewRouter(h)

	log.Println("Сервер запущен на порту :8080")
	return http.ListenAndServe(":8080", r)
}
