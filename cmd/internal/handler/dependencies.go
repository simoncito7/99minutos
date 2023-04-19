package handler

import (
	"context"
	"time"

	"github.com/99minutos/internal/repository"
)

type OrderRequest struct {
	ID       int    `json:"id"`
	ClientID string `json:"client_id"`

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

type Client struct {
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type LoginClientRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginClientResponse struct {
	AccessToken string `json:"access_token"`
	Client      Client `json:"client"`
}

type Service interface {
	CreateClient(ctx context.Context, client repository.Client) error
	GetClient(ctx context.Context, username string) (repository.Client, error)

	CreateOrder(ctx context.Context, order repository.Order) error
	InquireOrder(ctx context.Context, id int) (repository.Order, error)
	UpdateOrder(ctx context.Context, incomingOrder repository.Order) (bool, error)
	CancelOrder(ctx context.Context, id int) (bool, error)

	// extra
	GetAllOrders(ctx context.Context) ([]repository.Order, error)
}
