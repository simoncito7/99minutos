package repository

import "time"

type Client struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Token     string    `json:"token" db:"token"`
}

type Order struct {
	ID       int `json:"id" db:"id"`
	ClientID int `json:"client_id" db:"client_id"`

	OriginAddress    string `json:"last_name" db:"last_name"`
	OriginPostalCode string `json:"origin_postal_code" db:"origin_postal_code"`
	OriginExtNum     string `json:"origin_ext_num" db:"origin_ext_num"`
	OriginIntNum     string `json:"origin_int_num" db:"origin_int_num"`
	OriginCity       string `json:"origin_city" db:"origin_city"`

	DestinationAddress    string `json:"destination_address" db:"destination_address"`
	DestinationPostalCode string `json:"destination_postal_code" db:"destination_postal_code"`
	DestinationExtNum     string `json:"destination_ext_num" db:"destination_ext_num"`
	DestinationIntNum     string `json:"destination_int_num" db:"destination_int_num"`
	DestinationCity       string `json:"destination_city" db:"destination_city"`

	ProductQuantity int       `json:"product_quantity" db:"product_quantity"`
	TotalWeight     float64   `json:"total_weight" db:"total_weight"`
	PackageSize     string    `json:"package_size" db:"package_size"`
	Status          string    `json:"status" db:"status"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	WasRefunded     string    `json:"was_refunded" db:"was_refunded"`
}

type Auth struct {
	ID       int    `json:"id" db:"id"`
	ClientID int    `json:"client_id" db:"client_id"`
	Token    string `json:"token" db:"token"`
}
