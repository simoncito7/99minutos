package service

import (
	"context"

	"github.com/99minutos/internal/repository"
)

type Repository interface {
	CreateClient(ctx context.Context, client repository.Client) error
	GetClient(ctx context.Context, username string) (repository.Client, error)

	CreateOrder(ctx context.Context, order repository.Order) error
	GetOrder(ctx context.Context, id int) (repository.Order, error)
	UpdateOrderStatus(ctx context.Context, order repository.Order) error
	DeleteOrder(ctx context.Context, id int) error
	GetAllOrders(ctx context.Context) ([]repository.Order, error)
}
