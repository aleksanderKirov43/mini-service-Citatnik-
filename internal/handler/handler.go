package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mini-service-Citatnik-/internal/model"
	"mini-service-Citatnik-/internal/service"

	"github.com/gorilla/mux"
)

type QuoteHandler struct {
	service *service.QuoteService
}

func NewQuoteHandler() *QuoteHandler {
	return &QuoteHandler{
		service: service.NewQuoteService(),
	}
}

func (h *QuoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var q model.Quote
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "Недопустимое тело запроса", http.StatusBadRequest)
		return
	}

	if q.Author == "" || q.Quote == "" {
		http.Error(w, "Поля author и quote обязательны для заполнения", http.StatusBadRequest)
		return
	}

	created := h.service.AddQuote(&q)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *QuoteHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	quotes := h.service.GetAll(author)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	q, err := h.service.GetRandom()
	if err == service.ErrNotFound {
		http.Error(w, "Цитаты не найдены", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "id не существует", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err == service.ErrNotFound {
		http.Error(w, "Цитата не найдена", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Ошибка удаления цитаты", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
