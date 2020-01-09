package handler

import (
	"github.com/fwchen/jellyfish/application/middleware"
	"github.com/fwchen/jellyfish/domain/taco/repository"
	"github.com/fwchen/jellyfish/domain/taco/service"
	"github.com/juju/errors"
	"github.com/labstack/echo"
	"net/http"
)

type handler struct {
	tacoService *service.TacoApplicationService
}

func NewHandler(tacoRepo repository.Repository) *handler {
	return &handler{tacoService: service.NewTacoApplicationService(tacoRepo)}
}

func (h *handler) GetTacos(c echo.Context) error {
	userID := middleware.GetClaimsUserID(c)
	tacos, err := h.tacoService.GetTacos(userID, repository.ListTacoFilter{})
	if err != nil {
		return errors.Trace(err)
	}
	return c.JSON(http.StatusOK, tacos)
}