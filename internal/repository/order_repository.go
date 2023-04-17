package repository

import (
	"context"
	"fmt"
)

const (
	_queryCreateOrder = `INSERT INTO "order" (
        client_id,
        origin_address,
        origin_postal_code,
        origin_ext_num,
        origin_int_num,
        origin_city,
        destination_address,
        destination_postal_code,
        destination_ext_num,
        destination_int_num,
        destination_city,
        product_quantity,
        total_weight,
        package_size,
        status,
        created_at,
        updated_at,
        was_refunded
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)`

	_queryInquiryOrder = `SELECT * FROM "order" WHERE id = $1`
)

func (r *Repository) CreateOrder(ctx context.Context, order Order) error {
	_, err := r.db.Exec(
		_queryCreateOrder,
		order.ClientID,
		order.OriginAddress,
		order.OriginPostalCode,
		order.OriginExtNum,
		order.OriginIntNum,
		order.OriginCity,
		order.DestinationAddress,
		order.DestinationPostalCode,
		order.DestinationExtNum,
		order.DestinationIntNum,
		order.DestinationCity,
		order.ProductQuantity,
		order.TotalWeight,
		order.PackageSize,
		order.Status,
		order.CreatedAt,
		order.UpdatedAt,
		order.WasRefunded,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) InquireOrder(ctx context.Context, id int) (Order, error) {
	var order Order
	err := r.db.Get(&order, _queryInquiryOrder, id)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

func (r *Repository) UpdateOrderStatus(ctx context.Context, client Client) error {
	fmt.Println("Hello!")
	return nil
}

func (r *Repository) CancelOrder(ctx context.Context, client Client) error {
	return nil
}
