package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/99minutos/internal/repository"
	"github.com/99minutos/internal/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	// repo    *repository.Repository
	router  *gin.Engine
	service Service
}

func NewServer(serv *service.Service) *Server {
	server := &Server{service: serv}
	router := gin.Default()

	router.GET("/test")
	router.POST("/new-order", server.createOrder)
	router.GET("/order/:id", server.inquireOrder)
	router.GET("/orders", server.getAllOrders)
	router.PUT("/order-update", server.updateOrderStatus)
	router.DELETE("/order/:id", server.cancelOrder)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) createOrder(ctx *gin.Context) {
	var request OrderRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = server.service.CreateOrder(ctx, repository.Order(request))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (server *Server) inquireOrder(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	fmt.Printf("GetOrder called with id=%d\n", id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := server.service.InquireOrder(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (server *Server) updateOrderStatus(ctx *gin.Context) {
	var request OrderRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wasUpdated, err := server.service.UpdateOrder(ctx, repository.Order(request))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !wasUpdated {
		ctx.JSON(http.StatusOK, gin.H{"message": "there wasn't any update in the order status"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "order status updated successfully"})
}

func (server *Server) cancelOrder(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wasRefunded, err := server.service.CancelOrder(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("order was refunded? --> %v", wasRefunded))
}

func (server *Server) getAllOrders(ctx *gin.Context) {
	orders, err := server.service.GetAllOrders(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
