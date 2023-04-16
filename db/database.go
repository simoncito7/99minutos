package db

import (
	"context"
	"fmt"

	"github.com/99minutos/settings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.Port,
		s.DB.Name)

	return sqlx.ConnectContext(ctx, "mysql", connectionString)
}
