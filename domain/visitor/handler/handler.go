package handler

import (
	configs "github.com/fwchen/jellyfish/config"
	"github.com/fwchen/jellyfish/domain/visitor/repository"
	visitorService "github.com/fwchen/jellyfish/domain/visitor/service"
	"github.com/juju/errors"
	"github.com/labstack/echo"
	"net/http"
)

type handler struct {
	service *visitorService.ApplicationService
	config  *configs.ApplicationConfig
}

func NewHandler(visitorRepo repository.Repository, config *configs.ApplicationConfig) *handler {
	service := visitorService.NewApplicationService(visitorRepo, config)
	return &handler{service: service}
}

func (h *handler) Login(c echo.Context) error {
	request := new(struct {
		username string `json:"username" validate:"required"`
		password string `json:"username" validate:"required"`
	})
	c.Bind(request)
	err := h.service.Login(request.username, request.password)
	if err != nil {
		return errors.Trace(err)
	}
	return c.NoContent(http.StatusOK)
}

func (h *handler) SignUp(c echo.Context) error {
	request := new(struct {
		username string `json:"username" validate:"required"`
		password string `json:"username" validate:"required"`
	})
	c.Bind(request)
	token, err := h.service.SignUp(request.username, request.password)
	if err != nil {
		return errors.Trace(err)
	}
	c.Response().Header().Set(h.config.JwtHeaderKey, *token)
	return c.NoContent(http.StatusOK)
}
