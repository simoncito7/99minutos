package service

import (
	"context"

	"github.com/99minutos/internal/repository"
)

type Repository interface {
	CreateOrder(ctx context.Context, order repository.Order) error
	InquireOrder(ctx context.Context, id int) (repository.Order, error)
	UpdateOrderStatus(ctx context.Context, order repository.Order) error
	CancelOrder(ctx context.Context, client repository.Order) error
}
