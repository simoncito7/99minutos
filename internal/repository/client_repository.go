package repository

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

const (
	_queryCreateUser = `INSERT INTO clients (
		username,
		fullname,
		email,
		password,
		created_at,
		token
	  ) VALUES ($1, $2, $3, $4, $5, $6);
	  `

	_queryGetClient = `SELECT * FROM clients WHERE username = $1`
)

func (r *Repository) CreateClient(ctx context.Context, client Client) error {
	bytePass, err := hashPassword(client.Password)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, _queryCreateUser, client.Username, client.Fullname, client.Email, string(bytePass), client.CreatedAt, client.Token)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetClient(ctx context.Context, username string) (Client, error) {
	var client Client
	err := r.db.Get(&client, _queryGetClient, username)
	if err != nil {
		return Client{}, err
	}

	return client, nil
}

func hashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}
