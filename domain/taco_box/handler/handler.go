package handler

import (
	"jellyfish/application/middleware"
	"jellyfish/domain/taco_box/command"
	"jellyfish/domain/taco_box/repository"
	"jellyfish/domain/taco_box/service"
	"github.com/juju/errors"
	"github.com/labstack/echo"
	"net/http"
)

type handler struct {
	tacoBoxService *service.TacoBoxApplicationService
}

func NewHandler(tacoBoxRepo repository.TacoBoxRepository) *handler {
	return &handler{tacoBoxService: service.NewTacoBoxApplicationService(tacoBoxRepo)}
}

func (h *handler) GetTacoBoxes(c echo.Context) error {
	userID := middleware.GetClaimsUserID(c)
	tacos, err := h.tacoBoxService.GetTacoBoxes(userID)
	if err != nil {
		return errors.Trace(err)
	}
	return c.JSON(http.StatusOK, tacos)
}

func (h *handler) CreateTacoBox(c echo.Context) error {
	userID := middleware.GetClaimsUserID(c)
	var requestCommand command.CreateTacoBoxCommand
	requestCommand.CreatorId = userID
	err := c.Bind(&requestCommand)
	if err != nil {
		return errors.Trace(err)
	}
	_, err = h.tacoBoxService.CreateTacoBox(requestCommand)
	if err != nil {
		return errors.Trace(err)
	}
	return c.NoContent(http.StatusCreated)
}

func (h *handler) UpdateTacoBox(c echo.Context) error {
	userID := middleware.GetClaimsUserID(c)
	tacoID := c.Param("tacoID")
	var requestCommand command.UpdateTacoCommand
	requestCommand.TacoBoxID = tacoID
	requestCommand.OperationUserID = userID
	err := h.tacoBoxService.UpdateTacoBox(requestCommand)
	if err != nil {
		return errors.Trace(err)
	}
	return c.NoContent(http.StatusOK)
}
