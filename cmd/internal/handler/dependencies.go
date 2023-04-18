package handler

import (
	"context"
	"time"

	"github.com/99minutos/internal/repository"
)

type OrderRequest struct {
	ID       int `json:"id"`
	ClientID int `json:"client_id"`

	OriginAddress    string `json:"origin_address"`
	OriginPostalCode string `json:"origin_postal_code"`
	OriginExtNum     string `json:"origin_ext_num"`
	OriginIntNum     string `json:"origin_int_num"`
	OriginCity       string `json:"origin_city"`

	DestinationAddress    string `json:"destination_address"`
	DestinationPostalCode string `json:"destination_postal_code"`
	DestinationExtNum     string `json:"destination_ext_num"`
	DestinationIntNum     string `json:"destination_int_num"`
	DestinationCity       string `json:"destination_city"`

	ProductQuantity int       `json:"product_quantity"`
	TotalWeight     float64   `json:"total_weight"`
	PackageSize     string    `json:"package_size" validate:"oneof=S M L"`
	Status          string    `json:"status" validate:"oneof=creado recolectado en_estacion en_ruta entregado cancelado"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	WasRefunded     bool      `json:"was_refunded"`
}

type Service interface {
	CreateOrder(ctx context.Context, order repository.Order) error
	InquireOrder(ctx context.Context, id int) (repository.Order, error)
	UpdateOrder(ctx context.Context, incomingOrder repository.Order) (bool, error)
	CancelOrder(ctx context.Context, id int) (bool, error)

	// extra
	GetAllOrders(ctx context.Context) ([]repository.Order, error)
}
