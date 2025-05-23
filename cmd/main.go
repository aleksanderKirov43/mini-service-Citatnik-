package main

import (
	"log"

	"mini-service-Citatnik-/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
