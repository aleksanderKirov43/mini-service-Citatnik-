package service

import (
	"context"
	"testing"

	"mini-service-Citatnik-/internal/model"
)

func TestAddQuote(t *testing.T) {
	s := NewQuoteService()
	ctx := context.Background()
	q := &model.Quote{
		Author: "Тестовый автор",
		Quote:  "А это его тестовая цитата",
	}

	created := s.AddQuote(ctx, q)
	if created.ID != 1 {
		t.Errorf("Ожидаемый ID: 1, получено %d", created.ID)
	}
}

func TestGetAllQuotes(t *testing.T) {
	s := NewQuoteService()
	ctx := context.Background()
	s.AddQuote(ctx, &model.Quote{Author: "Пушкин", Quote: "Тест А"})
	s.AddQuote(ctx, &model.Quote{Author: "Лермонтов", Quote: "Тест Б"})

	all := s.GetAll(ctx, "")
	if len(all) != 2 {
		t.Errorf("Ожидалось количество цитат: 2, получено %d", len(all))
	}

	filtered := s.GetAll(ctx, "Пушкин")
	if len(filtered) != 1 || filtered[0].Author != "Пушкин" {
		t.Errorf("Ожидалась одна цитата Пушкина, получено %d", len(filtered))
	}
}

func TestGetRandomQuote(t *testing.T) {
	s := NewQuoteService()
	ctx := context.Background()
	_, err := s.GetRandom(ctx)
	if err == nil {
		t.Errorf("Хранилище пустое!")
	}

	s.AddQuote(ctx, &model.Quote{Author: "X", Quote: "Y"})
	q, err := s.GetRandom(ctx)
	if err != nil || q == nil {
		t.Errorf("Ошибка получения цитаты")
	}
}

func TestDeleteQuote(t *testing.T) {
	s := NewQuoteService()
	ctx := context.Background()
	q := s.AddQuote(ctx, &model.Quote{Author: "Del", Quote: "Z"})

	err := s.Delete(ctx, int64(q.ID))
	if err != nil {
		t.Errorf("Ошибка при удалении")
	}

	err = s.Delete(ctx, 999)
	if err != ErrNotFound {
		t.Errorf("Ошибка: не найдено, получено %v", err)
	}
}
