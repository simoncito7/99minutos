package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/99minutos/internal/repository"
	"github.com/labstack/gommon/log"
)

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateOrder(ctx context.Context, order repository.Order) error {
	_, err := s.repo.GetClient(ctx, order.ClientID)
	if err != nil {
		return err
	}

	if !s.checkAllowedWeightAndSize(order.PackageSize, order.TotalWeight) {
		if order.TotalWeight > 25 {
			return errors.New("for package weight bigger than 25kg, please contact us")
		}
		return errors.New("package size or weight not allowed")
	}

	order.CreatedAt = time.Now()
	err = s.repo.CreateOrder(ctx, order)
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

	now := time.Now()

	refund := s.wasOrderCancelledBeforeTwoMinutes(storedOrder.CreatedAt, now)
	log.Debug(ctx, storedOrder.CreatedAt, now)
	log.Info(ctx, storedOrder.CreatedAt, now)
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

func (s *Service) CreateClient(ctx context.Context, client repository.Client) error {
	err := s.repo.CreateClient(ctx, client)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetClient(ctx context.Context, username string) (repository.Client, error) {
	client, err := s.repo.GetClient(ctx, username)
	if err != nil {
		return repository.Client{}, err
	}

	return client, nil
}

func (s *Service) wasOrderCancelledBeforeTwoMinutes(created, updated time.Time) bool {
	diff := updated.Sub(created)

	return diff.Minutes() < 2
}

// checkAllowedWeightAndSize checks if the order is the weight is allowed for the
// corresponding size.
func (s *Service) checkAllowedWeightAndSize(size string, weight float64) bool {
	switch size {
	case "S":
		return weight <= 5.0
	case "M":
		return weight > 5 && weight <= 15.0
	case "L":
		return weight > 15.0 && weight <= 25.0
	}

	return false
}
