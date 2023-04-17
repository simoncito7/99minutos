package handler

import (
	"github.com/99minutos/internal/service"
	"github.com/labstack/echo/v4"
)

type API struct {
	service service.Service
}

func New(serv service.Service) *API {
	return &API{service: serv}
}

func (a *API) Routes(e *echo.Echo) {
}
