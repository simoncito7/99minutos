package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/99minutos/internal/repository"
	"github.com/99minutos/internal/service"
	"github.com/99minutos/settings"
	"github.com/99minutos/token"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Println(server.cfg)        // add this line
	fmt.Println(server.tokenMaker) // add this line
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
