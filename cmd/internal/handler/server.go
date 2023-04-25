package handler

import (
	"fmt"

	"github.com/99minutos/internal/service"
	"github.com/99minutos/settings"
	"github.com/99minutos/token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	service    Service
	tokenMaker token.Maker
	cfg        *settings.Settings
}

func NewServer(serv *service.Service, cfg *settings.Settings) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(cfg.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}
	router := gin.Default()
	server := &Server{service: serv, cfg: cfg, tokenMaker: tokenMaker}

	routes(router, server)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
