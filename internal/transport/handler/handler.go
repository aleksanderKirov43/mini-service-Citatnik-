package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mini-service-Citatnik-/internal/model"
	"mini-service-Citatnik-/internal/service"

	"github.com/gorilla/mux"
)

type IQuoteHandler interface {
	CreateQuote(w http.ResponseWriter, r *http.Request)
	GetQuotes(w http.ResponseWriter, r *http.Request)
	GetRandomQuote(w http.ResponseWriter, r *http.Request)
	DeleteQuote(w http.ResponseWriter, r *http.Request)
}

type QuoteHandler struct {
	service service.IQuoteService
}

func NewQuoteHandler(s service.IQuoteService) IQuoteHandler {
	return &QuoteHandler{
		service: s,
	}
}

func (h *QuoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var q model.Quote
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "Недопустимое тело запроса", http.StatusBadRequest)
		return
	}

	if q.Author == "" || q.Quote == "" {
		http.Error(w, "Поля author и quote обязательны для заполнения", http.StatusBadRequest)
		return
	}

	created := h.service.AddQuote(ctx, &q)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *QuoteHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	author := r.URL.Query().Get("author")
	quotes := h.service.GetAll(ctx, author)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	q, err := h.service.GetRandom(ctx)
	if err == service.ErrNotFound {
		http.Error(w, "Цитаты не найдены", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "id не существует", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(ctx, id)
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
