package service

import (
	"context"
	"errors"
	"strings"
	"time"

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

func (s *Service) InquireOrder(ctx context.Context, id int) (repository.Order, error) {
	order, err := s.repo.GetOrder(ctx, id)
	if err != nil {
		return repository.Order{}, err
	}

	return order, nil
}

func (s *Service) UpdateOrder(ctx context.Context, incomingOrder repository.Order) (bool, error) {
	storedOrder, err := s.repo.GetOrder(ctx, incomingOrder.ID)
	if err != nil {
		return false, err
	}

	if storedOrder.Status == incomingOrder.Status {
		return false, nil
	}

	err = s.repo.UpdateOrderStatus(ctx, incomingOrder)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *Service) CancelOrder(ctx context.Context, id int) (bool, error) {
	storedOrder, err := s.repo.GetOrder(ctx, id)
	if err != nil {
		return false, err
	}

	refund := s.wasOrderCancelledBeforeTwoMinutes(storedOrder.CreatedAt, time.Now())
	switch true {
	case strings.Contains(storedOrder.Status, "en_ruta"):
		return false, errors.New("status en_ruta: order cannot be cancelled in this status")
	case strings.Contains(storedOrder.Status, "entregado"):
		return false, errors.New("status entregado: order cannot be cancelled in this status")
	default:
		err := s.repo.DeleteOrder(ctx, id)
		if err != nil {
			return false, err
		}
	}

	return refund, nil
}

func (s *Service) GetAllOrders(ctx context.Context) ([]repository.Order, error) {
	orders, err := s.repo.GetAllOrders(ctx)
	if err != nil {
		return []repository.Order{}, err
	}

	return orders, nil
}

func (s *Service) wasOrderCancelledBeforeTwoMinutes(created, updated time.Time) bool {
	diff := updated.Sub(created)

	return diff.Minutes() < 2
}
