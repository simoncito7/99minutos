package handler

import (
	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine, server *Server) {
	router.GET("/test")
	router.POST("/client", server.createClient)
	router.POST("/client/login", server.loginClient)
	router.POST("/order", server.createOrder)
	router.GET("/order/:id", server.inquireOrder)
	router.GET("/orders", server.getAllOrders)
	router.PUT("/order/update", server.updateOrderStatus)
	router.DELETE("/order/:id", server.cancelOrder)
}
