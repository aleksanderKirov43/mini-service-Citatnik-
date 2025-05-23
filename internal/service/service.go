package service

import (
	"context"
	"errors"
	"math/rand"
	"mini-service-Citatnik-/internal/model"
	"strings"
	"sync"
	"time"
)

var (
	ErrNotFound = errors.New("quote not found")
)

type QuoteService struct {
	mu     sync.RWMutex
	quotes map[int]*model.Quote
	nextID int
}

type IQuoteService interface {
	AddQuote(ctx context.Context, q *model.Quote) *model.Quote
	GetAll(ctx context.Context, author string) []*model.Quote
	GetRandom(ctx context.Context) (*model.Quote, error)
	Delete(ctx context.Context, id int64) error
}

func NewQuoteService() IQuoteService {
	return &QuoteService{
		quotes: make(map[int]*model.Quote),
		nextID: 1,
	}
}

func (s *QuoteService) AddQuote(ctx context.Context, q *model.Quote) *model.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	q.ID = s.nextID
	s.nextID++
	s.quotes[q.ID] = q

	return q
}

func (s *QuoteService) GetAll(ctx context.Context, author string) []*model.Quote {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var res []*model.Quote
	author = strings.ToLower(author)

	for _, q := range s.quotes {
		if author == "" {
			res = append(res, q)
			continue
		}

		if strings.Contains(strings.ToLower(q.Author), author) {
			res = append(res, q)
		}
	}

	return res
}

func (s *QuoteService) GetRandom(ctx context.Context) (*model.Quote, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.quotes) == 0 {
		return nil, ErrNotFound
	}

	var allQuotes []*model.Quote
	for _, q := range s.quotes {
		allQuotes = append(allQuotes, q)
	}

	rand.Seed(time.Now().UnixNano())
	return allQuotes[rand.Intn(len(allQuotes))], nil
}

func (s *QuoteService) Delete(ctx context.Context, id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.quotes[int(id)]; !ok {
		return ErrNotFound
	}
	delete(s.quotes, int((id)))
	return nil
}
