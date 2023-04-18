package mocktest

import (
	"context"
	"testing"

	"github.com/99minutos/internal/repository"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
	t *testing.T
}

func NewRepositoryMock(t *testing.T) *RepositoryMock {
	return &RepositoryMock{t: t}
}

func (repo *RepositoryMock) CreateOrder(ctx context.Context, order repository.Order) error {
	args := repo.Called(order)
	return args.Error(0)
}

func (repo *RepositoryMock) GetOrder(ctx context.Context, id int) (repository.Order, error) {
	args := repo.Called(id)
	return args.Get(0).(repository.Order), args.Error(1)
}

func (repo *RepositoryMock) UpdateOrderStatus(ctx context.Context, order repository.Order) error {
	args := repo.Called(order)
	return args.Error(0)
}

func (repo *RepositoryMock) DeleteOrder(ctx context.Context, id int) error {
	args := repo.Called(id)
	return args.Error(0)
}

func (repo *RepositoryMock) GetAllOrders(ctx context.Context) ([]repository.Order, error) {
	args := repo.Called()
	return args.Get(0).([]repository.Order), args.Error(1)
}
