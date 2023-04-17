package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

const (
	_queryCreateUser = `INSERT INTO client (
		name,
		last_name,
		email,
		password,
		created_at,
		token
	  ) VALUES ($1, $2, $3, $4, $5, $6);
	  `

	_queryGetClient = `SELECT * FROM client WHERE id = $1`
)

func (r *Repository) CreateClient(ctx context.Context, db *sqlx.DB, client Client) error {
	_, err := db.ExecContext(ctx, _queryCreateUser, client.Name, client.LastName, client.Email, client.Password, client.CreatedAt, client.Token)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetClient(ctx context.Context, id int) (Client, error) {
	var client Client
	err := r.db.Get(&client, _queryGetClient, id)
	if err != nil {
		return Client{}, err
	}

	return client, nil
}
