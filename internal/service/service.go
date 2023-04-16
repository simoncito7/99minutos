package service

import (
	"context"

	"github.com/99minutos/internal/repository"
)

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateOrder(ctx context.Context, order repository.Order) error {
	err := s.repo.CreateOrder(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetOrder(ctx context.Context, id int) (repository.Order, error) {
	order, err := s.repo.InquireOrder(ctx, id)
	if err != nil {
		return repository.Order{}, err
	}

	return order, nil
}
