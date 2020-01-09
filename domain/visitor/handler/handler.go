package handler

import (
	"github.com/dchest/captcha"
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
	return &handler{service: service, config: config}
}

func (h *handler) Login(c echo.Context) error {
	request := new(struct {
		Username  string `json:"username" validate:"required"`
		Password  string `json:"password" validate:"required"`
		Captcha   string `json:"captcha" validate:"required"`
		CaptchaID string `json:"captchaID" validate:"required"`
	})
	c.Bind(request)
	if !captcha.VerifyString(request.CaptchaID, request.Captcha) {
		return c.NoContent(http.StatusBadRequest)
	}
	token, err := h.service.Login(request.Username, request.Password)
	if err != nil {
		return errors.Trace(err)
	}
	if token == nil {
		return c.NoContent(http.StatusUnauthorized)
	}
	c.Response().Header().Set(h.config.JwtHeaderKey, *token)
	return c.NoContent(http.StatusOK)
}

func (h *handler) SignUp(c echo.Context) error {
	request := new(struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	})
	c.Bind(request)
	err := h.service.SignUp(request.Username, request.Password)
	if err != nil {
		return errors.Trace(err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *handler) GenCaptcha(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{
		"id": captcha.New(),
	})
}
