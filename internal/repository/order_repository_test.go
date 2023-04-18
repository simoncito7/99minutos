package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRepository_CreateOrder(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	repository := New(db)

	newOrder := createFakeOrder()

	err := repository.CreateOrder(context.Background(), newOrder)
	require.NoError(t, err)

	var order Order
	err = db.Get(&order, fmt.Sprintf(`SELECT * FROM "order" WHERE client_id = $%d`, newOrder.ClientID), newOrder.ClientID)
	require.NoError(t, err)
	require.Equal(t, newOrder.ClientID, order.ClientID)
	require.Equal(t, newOrder.OriginAddress, order.OriginAddress)
	require.Equal(t, newOrder.OriginPostalCode, order.OriginPostalCode)
	require.Equal(t, newOrder.OriginExtNum, order.OriginExtNum)
	require.Equal(t, newOrder.OriginIntNum, order.OriginIntNum)
	require.Equal(t, newOrder.OriginCity, order.OriginCity)
	require.Equal(t, newOrder.DestinationAddress, order.DestinationAddress)
	require.Equal(t, newOrder.DestinationPostalCode, order.DestinationPostalCode)
	require.Equal(t, newOrder.DestinationExtNum, order.DestinationExtNum)
	require.Equal(t, newOrder.DestinationIntNum, order.DestinationIntNum)
	require.Equal(t, newOrder.DestinationCity, order.DestinationCity)
	require.Equal(t, newOrder.ProductQuantity, order.ProductQuantity)
	require.Equal(t, newOrder.TotalWeight, order.TotalWeight)
	require.Equal(t, newOrder.PackageSize, order.PackageSize)
	require.Equal(t, newOrder.Status, order.Status)
	require.Equal(t, newOrder.WasRefunded, order.WasRefunded)
}

func TestInquireOrder(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	repo := New(db)

	order := createFakeOrder()

	// In this case I will check with an existent element just for testing
	// I already checked that for id = 4 there is an stored element in the table "order"

	// retrieve the order by its ID
	retrievedOrder, err := repo.GetOrder(context.Background(), 4)
	require.NoError(t, err)

	require.Equal(t, retrievedOrder.ClientID, order.ClientID)
	require.Equal(t, retrievedOrder.OriginAddress, order.OriginAddress)
	require.Equal(t, retrievedOrder.OriginPostalCode, order.OriginPostalCode)
	require.Equal(t, retrievedOrder.OriginExtNum, order.OriginExtNum)
	require.Equal(t, retrievedOrder.OriginIntNum, order.OriginIntNum)
	require.Equal(t, retrievedOrder.OriginCity, order.OriginCity)
	require.Equal(t, retrievedOrder.DestinationAddress, order.DestinationAddress)
	require.Equal(t, retrievedOrder.DestinationPostalCode, order.DestinationPostalCode)
	require.Equal(t, retrievedOrder.DestinationExtNum, order.DestinationExtNum)
	require.Equal(t, retrievedOrder.DestinationIntNum, order.DestinationIntNum)
	require.Equal(t, retrievedOrder.DestinationCity, order.DestinationCity)
	require.Equal(t, retrievedOrder.ProductQuantity, order.ProductQuantity)
	require.Equal(t, retrievedOrder.TotalWeight, order.TotalWeight)
	require.Equal(t, retrievedOrder.PackageSize, order.PackageSize)
	require.Equal(t, retrievedOrder.Status, order.Status)

	// retrieve the order by its ID
	retrievedOrder, err = repo.GetOrder(context.Background(), 1)
	require.Error(t, err)
	require.Empty(t, retrievedOrder)
}

func createFakeOrder() Order {
	return Order{
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

func TestUpdateOrder(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	repo := New(db)

	order := createFakeOrder()

	err := repo.CreateOrder(context.Background(), order)
	require.NoError(t, err)

	// order status updated to "en_ruta".
	updatedOrder := Order{
		ID:        6,
		Status:    "en_ruta",
		UpdatedAt: time.Now(),
	}

	err = repo.UpdateOrderStatus(context.Background(), updatedOrder)
	require.NoError(t, err)

	// here we retrieve the order from the database and verify that
	// its status has been updated.
	retrievedOrder, err := repo.GetOrder(context.Background(), 6)
	require.NoError(t, err)
	require.Equal(t, "en_ruta", retrievedOrder.Status)
}

func TestDeleteOrder(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	repo := New(db)

	order := createFakeOrder()

	err := repo.CreateOrder(context.Background(), order)
	require.NoError(t, err)

	err = repo.DeleteOrder(context.Background(), order.ID)
	require.NoError(t, err)

	// retrieve the order from the database to check if it was deleted
	retrievedOrder, err := repo.GetOrder(context.Background(), order.ID)
	require.Error(t, err)
	require.Equal(t, Order{}, retrievedOrder)
}
