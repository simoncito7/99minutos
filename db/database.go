package db

import (
	"context"
	"fmt"

	"github.com/99minutos/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.DB.Host, s.DB.Port, s.DB.User, s.DB.Password, s.DB.Name)

	return sqlx.ConnectContext(ctx, "postgres", connectionString)
}
