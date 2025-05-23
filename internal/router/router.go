package router

import (
	"net/http"

	"mini-service-Citatnik-/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {

	r := mux.NewRouter()
	h := handler.NewQuoteHandler()

	r.HandleFunc("/quotes", h.CreateQuote).Methods(http.MethodPost)
	r.HandleFunc("/quotes", h.GetQuotes).Methods(http.MethodGet)
	r.HandleFunc("/quotes/random", h.GetRandomQuote).Methods(http.MethodGet)
	r.HandleFunc("/quotes/{id}", h.DeleteQuote).Methods(http.MethodDelete)

	return r
}
