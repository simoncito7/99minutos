package handler

import (
	"github.com/99minutos/cmd/internal/authentication"
	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine, server *Server) {
	router.GET("/test")
	router.POST("/client", server.createClient)
	router.POST("/client/login", server.loginClient)

	authRoutes := router.Group("/").Use(authentication.AuthMiddelware(server.tokenMaker))
	authRoutes.POST("/order", server.createOrder)
	authRoutes.GET("/order/:id", server.inquireOrder)
	authRoutes.GET("/orders", server.getAllOrders)
	authRoutes.PUT("/order/update", server.updateOrderStatus)
	authRoutes.DELETE("/order/:id", server.cancelOrder)
}
