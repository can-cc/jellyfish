package handler

import (
	"net/http"

	"jellyfish/application/middleware"
	"jellyfish/domain/taco"
	tacoCommand "jellyfish/domain/taco/command"
	"jellyfish/domain/taco/repository"
	"jellyfish/domain/taco/service"
	"jellyfish/domain/taco_box"
	boxService "jellyfish/domain/taco_box/service"
	"github.com/juju/errors"
	"github.com/labstack/echo"
)

type handler struct {
	tacoService *service.TacoApplicationService
}

func NewHandler(tacoRepo repository.Repository, tacoBoxPermissionService *boxService.TacoBoxPermissionService) *handler {
	return &handler{tacoService: service.NewTacoApplicationService(tacoRepo, tacoBoxPermissionService)}
}

func (h *handler) GetTacos(c echo.Context) error {
	userID := middleware.GetClaimsUserID(c)
	statues := taco.ParseStatues(c.QueryParam("status"))
	qtype := c.QueryParam("type")
	boxId := c.QueryParam("boxId")
	isImportant := c.QueryParam("isImportant") != "" && c.QueryParam("isImportant") != "false"
	isScheduled := c.QueryParam("isScheduled") != "" && c.QueryParam("isScheduled") != "false"

	var tacoType *taco.Type
	if qtype == "" {
		tacoType = nil
	} else {
		t := taco.Type(qtype)
		tacoType = &t
	}

	filter := taco.TacoFilter{
		BoxId:     taco_box.PointerBoxIdIfEmptyStr(boxId),
		Type:      tacoType,
		Statues:   statues,
		Important: isImportant,
		Scheduled: isScheduled,
	}

	tacos, err := h.tacoService.GetTacos(userID, filter)
	if err != nil {
		return errors.Trace(err)
	}
	return c.JSON(http.StatusOK, tacos)
}

func (h *handler) CreateTaco(c echo.Context) error {
	userID := middleware.GetClaimsUserID(c)
	var command tacoCommand.CreateTacoCommand
	err := c.Bind(&command)
	if err != nil {
		return errors.Trace(err)
	}
	_, err = h.tacoService.CreateTaco(&command, userID)
	if err != nil {
		return errors.Trace(err)
	}
	return c.NoContent(http.StatusCreated)
}

func (h *handler) UpdateTaco(c echo.Context) error {
	userID := middleware.GetClaimsUserID(c)
	tacoID := c.Param("tacoId")
	var command tacoCommand.UpdateTacoCommand
	err := c.Bind(&command)
	if err != nil {
		return errors.Trace(err)
	}
	command.TacoId = tacoID
	command.OperationUserID = userID
	err = h.tacoService.UpdateTaco(command)
	if err != nil {
		return errors.Trace(err)
	}
	return c.NoContent(http.StatusOK)
}

func (h *handler) DeleteTaco(c echo.Context) error {
	tacoId := c.Param("tacoId")
	err := h.tacoService.DeleteTaco(tacoId)
	if err != nil {
		return errors.Trace(err)
	}
	return c.NoContent(http.StatusOK)
}

func (h *handler) SortTaco(c echo.Context) error {
	userID := middleware.GetClaimsUserID(c)
	var command tacoCommand.SortTacoCommand
	err := c.Bind(&command)
	if err != nil {
		return errors.Trace(err)
	}
	err = h.tacoService.Reorder(&command, userID)
	if err != nil {
		return errors.Trace(err)
	}
	return c.NoContent(http.StatusOK)
}
