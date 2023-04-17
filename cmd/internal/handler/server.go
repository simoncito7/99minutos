package handler

import (
	"net/http"

	"github.com/99minutos/internal/repository"
	"github.com/gin-gonic/gin"
)

type Server struct {
	repo   *repository.Repository
	router *gin.Engine
}

func NewServer(repo *repository.Repository) *Server {
	server := &Server{repo: repo}
	router := gin.Default()

	router.GET("/test")
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) Hi() {
	var ctx gin.Context
	ctx.JSON(http.StatusOK, "hey!")
}
