package main

import (
	"context"
	"fmt"
	"log"

	"github.com/99minutos/db"
	"github.com/99minutos/internal/repository"
	"github.com/99minutos/internal/service"
	"github.com/99minutos/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.NewSettings,
			db.New,
			repository.New,
			service.New,
		),

		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("select * from clients")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	fmt.Println("Hello, world!")

	app.Run()

	return nil
}
