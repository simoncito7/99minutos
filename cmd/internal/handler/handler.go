package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/99minutos/cmd/internal/authentication"
	"github.com/99minutos/internal/repository"
	"github.com/99minutos/token"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) createOrder(ctx *gin.Context) {
	var request OrderRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authPayload := ctx.MustGet(authentication.AuthPayloadKey).(*token.Payload)

	request.ClientID = authPayload.Username
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
		ctx.JSON(http.StatusOK, gin.H{"message": "there are not any updates for this order"})
		return
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

func (server *Server) createClient(ctx *gin.Context) {
	var request Client
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = server.service.CreateClient(ctx, repository.Client{
		Username:  request.Username,
		Fullname:  request.Fullname,
		Password:  request.Password,
		Email:     request.Email,
		CreatedAt: request.CreatedAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (server *Server) loginClient(ctx *gin.Context) {
	var request LoginClientRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := server.service.GetClient(ctx, request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !checkPasswork(client.Password, request.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(client.Username, server.cfg.TokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	loginResponse := LoginClientResponse{
		AccessToken: accessToken,
		Client: Client{
			Username:  client.Username,
			Fullname:  client.Fullname,
			Email:     client.Email,
			CreatedAt: client.CreatedAt,
		},
	}

	ctx.JSON(http.StatusOK, loginResponse)
}

func (server *Server) getAllOrders(ctx *gin.Context) {
	orders, err := server.service.GetAllOrders(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func checkPasswork(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
