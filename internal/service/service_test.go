package service

import (
	"testing"

	"mini-service-Citatnik-/internal/model"
)

func TestAddQuote(t *testing.T) {
	s := NewQuoteService()
	q := &model.Quote{
		Author: "Тестовый автор",
		Quote:  "А это его тестовая цитата",
	}

	created := s.AddQuote(q)
	if created.ID != 1 {
		t.Errorf("Ожидаемый ID: 1, получено %d", created.ID)
	}
}

func TestGetAllQuotes(t *testing.T) {
	s := NewQuoteService()
	s.AddQuote(&model.Quote{Author: "Пушкин", Quote: "Тест А"})
	s.AddQuote(&model.Quote{Author: "Лермонтов", Quote: "Тест Б"})

	all := s.GetAll("")
	if len(all) != 2 {
		t.Errorf("Ожидалось количество цитат: 2, получено %d", len(all))
	}

	filtered := s.GetAll("Пушкин")
	if len(filtered) != 1 || filtered[0].Author != "Пушкин" {
		t.Errorf("Ожидалась одна цитата Пушкина, получено %d", len(filtered))
	}
}

func TestGetRandomQuote(t *testing.T) {
	s := NewQuoteService()
	_, err := s.GetRandom()
	if err == nil {
		t.Errorf("Хранилище пустое!")
	}

	s.AddQuote(&model.Quote{Author: "X", Quote: "Y"})
	q, err := s.GetRandom()
	if err != nil || q == nil {
		t.Errorf("Ошибка получения цитаты")
	}
}

func TestDeleteQuote(t *testing.T) {
	s := NewQuoteService()
	q := s.AddQuote(&model.Quote{Author: "Del", Quote: "Z"})

	err := s.Delete(int64(q.ID))
	if err != nil {
		t.Errorf("Ошибка при удалении")
	}

	err = s.Delete(999)
	if err != ErrNotFound {
		t.Errorf("Ошибка: не найдено, получено %v", err)
	}
}
