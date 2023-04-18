package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/99minutos/internal/repository"
	"github.com/99minutos/test/mocktest"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	repoMock := mocktest.NewRepositoryMock(t)

	order := getFakeOrder()
	repoMock.On("CreateOrder", order).Return(nil)

	service := New(repoMock)
	err := service.CreateOrder(context.Background(), order)
	require.NoError(t, err)
	repoMock.AssertExpectations(t)
}

func TestGetOrder(t *testing.T) {
	repoMock := mocktest.NewRepositoryMock(t)

	order := repository.Order{
		ID: 4,
	}
	repoMock.On("GetOrder", order.ID).Return(order, nil)

	service := New(repoMock)
	result, err := service.InquireOrder(context.Background(), order.ID)
	require.NoError(t, err)
	require.Equal(t, order, result)
	repoMock.AssertExpectations(t)
}

func TestUpdateOrder_StatusUpdatedSuccessfully(t *testing.T) {
	repoMock := mocktest.NewRepositoryMock(t)

	storedOrder := repository.Order{
		ID:     1,
		Status: "creado",
	}
	repoMock.On("GetOrder", storedOrder.ID).Return(storedOrder, nil)

	incomingOrder := repository.Order{
		ID:     1,
		Status: "en_ruta",
	}
	repoMock.On("UpdateOrderStatus", incomingOrder).Return(nil)

	service := New(repoMock)
	updated, err := service.UpdateOrder(context.Background(), incomingOrder)
	require.NoError(t, err)
	require.True(t, updated)
}

func TestUpdateOrder_NoUpdate(t *testing.T) {
	repoMock := mocktest.NewRepositoryMock(t)
	service := New(repoMock)

	storedOrder := repository.Order{
		ID:     1,
		Status: "en_ruta",
	}
	repoMock.On("GetOrder", storedOrder.ID).Return(storedOrder, nil)

	incomingOrder := repository.Order{
		ID:     1,
		Status: "en_ruta",
	}

	updated, err := service.UpdateOrder(context.Background(), incomingOrder)
	require.NoError(t, err)
	require.False(t, updated)
}

func TestUpdateOrder_OrderDoesntExist(t *testing.T) {
	_errNotFound := errors.New("not found")
	repoMock := mocktest.NewRepositoryMock(t)
	service := New(repoMock)

	orderID := 1
	repoMock.On("GetOrder", orderID).Return(repository.Order{}, _errNotFound)

	incomingOrder := repository.Order{
		ID:     orderID,
		Status: "en_ruta",
	}

	updated, err := service.UpdateOrder(context.Background(), incomingOrder)
	require.ErrorIs(t, err, _errNotFound)
	require.False(t, updated)
}

func TestCancelOrderCases(t *testing.T) {
	orders := []repository.Order{
		{
			ID:     1,
			Status: "pendiente",
		},
		{
			ID:     2,
			Status: "en_ruta",
		},
		{
			ID:     3,
			Status: "entregado",
		},
		{
			ID:     4,
			Status: "pendiente",
		},
		{
			ID:        5,
			Status:    "creado",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now().Add(time.Second * 30),
		},
	}

	testCases := []struct {
		orderID    int
		wantRefund bool
		wantErr    error
	}{
		{
			orderID:    orders[1].ID,
			wantRefund: false,
			wantErr:    errors.New("status en_ruta: order cannot be cancelled in this status"),
		},
		{
			orderID:    orders[2].ID,
			wantRefund: false,
			wantErr:    errors.New("status entregado: order cannot be cancelled in this status"),
		},
		{
			orderID:    orders[0].ID,
			wantRefund: false,
			wantErr:    nil,
		},
		{
			orderID:    orders[3].ID,
			wantRefund: false,
			wantErr:    nil,
		},
		{
			orderID:    orders[4].ID,
			wantRefund: true,
			wantErr:    nil,
		},
		{
			orderID:    6,
			wantRefund: false,
			wantErr:    errors.New("order not found"),
		},
	}

	repoMock := mocktest.NewRepositoryMock(t)

	for _, tc := range testCases {
		service := New(repoMock)

		for i := range orders {
			if orders[i].ID == tc.orderID {
				order := orders[i]
				repoMock.On("GetOrder", order.ID).Return(order, nil)

				if tc.wantErr == nil {
					repoMock.On("DeleteOrder", order.ID).Return(nil)
				}

				wasRefunded, err := service.CancelOrder(context.Background(), order.ID)
				require.Equal(t, tc.wantRefund, wasRefunded)
				require.Equal(t, tc.wantErr, err)
				break
			}
		}
	}

	repoMock.AssertExpectations(t)
}

func getFakeOrder() repository.Order {
	return repository.Order{
		ClientID:              1,
		OriginAddress:         "origin address",
		OriginPostalCode:      "12345",
		OriginExtNum:          "1A",
		OriginIntNum:          "10",
		OriginCity:            "origin city",
		DestinationAddress:    "dest address",
		DestinationPostalCode: "67890",
		DestinationExtNum:     "2B",
		DestinationIntNum:     "20",
		DestinationCity:       "dest city",
		ProductQuantity:       2,
		TotalWeight:           1.5,
		PackageSize:           "S",
		Status:                "creado",
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
		WasRefunded:           false,
	}
}
