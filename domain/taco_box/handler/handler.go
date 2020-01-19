package handler

import (
	"github.com/fwchen/jellyfish/application/middleware"
	command "github.com/fwchen/jellyfish/domain/taco_box/command"
	"github.com/fwchen/jellyfish/domain/taco_box/repository"
	service "github.com/fwchen/jellyfish/domain/taco_box/service"
	"github.com/juju/errors"
	"github.com/labstack/echo"
	"net/http"
)

type handler struct {
	tacoBoxService *service.TacoBoxApplicationService
}

func NewHandler(tacoRepo repository.TacoBoxRepository) *handler {
	return &handler{tacoBoxService: service.NewTacoBoxApplicationService(tacoRepo)}
}

func (h *handler) CreateTacoBox(c echo.Context) error {
	userID := middleware.GetClaimsUserID(c)
	var requestCommand command.CreateTacoBoxCommand
	requestCommand.CreatorId = userID
	err := c.Bind(&c)
	if err != nil {
		return errors.Trace(err)
	}
	_, err = h.tacoBoxService.CreateTacoBox(requestCommand)
	if err != nil {
		return errors.Trace(err)
	}
	return c.NoContent(http.StatusCreated)
}
