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
	cfg, err := settings.LoadConfig()
	if err != nil {
		log.Fatal("problem loading config")
	}

	db, err := db.New(ctx, cfg)
	if err != nil {
		log.Fatal("problem with db", err)
	}

	repo := repository.New(db)

	serv := service.New(repo)

	server, err := handler.NewServer(serv, cfg)
	if err != nil {
		log.Fatal("can't initialize server")
	}

	err = server.Start(cfg.Address)
	if err != nil {
		log.Fatal("can't start server")
	}
}
