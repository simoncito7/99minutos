package main

import (
	"context"
	"log"

	"github.com/99minutos/cmd/internal/handler"
	"github.com/99minutos/db"
	"github.com/99minutos/internal/repository"
	"github.com/99minutos/internal/service"
	"github.com/99minutos/settings"
)

func main() {
	ctx := context.Background()
	s, err := settings.NewSettings()
	if err != nil {
		log.Fatal("problem with settings")
	}
	db, err := db.New(ctx, s)
	if err != nil {
		log.Fatal("problem with db", err)
	}

	repo := repository.New(db)

	serv := service.New(repo)

	server := handler.NewServer(serv)

	err = server.Start("127.0.0.1:8080")
	if err != nil {
		log.Fatal("can't start server")
	}
}
