package handler

import (
	"github.com/fwchen/jellyfish/application"
	"github.com/fwchen/jellyfish/domain/user"
	"github.com/fwchen/jellyfish/domain/user/repository"
	"github.com/juju/errors"
	"github.com/labstack/echo"
	"net/http"
)

type handler struct {
	service *user.ApplicationService
}

func NewHandler(userRepo repository.Repository) *handler {
	return &handler{service: user.NewApplicationService(userRepo)}
}

func (h *handler) GetUserInfo(c echo.Context) error {
	userID := c.Param("userID")
	userInfo, err := h.service.GetUserInfo(userID)
	if err != nil {
		return errors.Trace(err)
	}
	return c.JSON(http.StatusOK, userInfo)
}

func (h *handler) UpdateUserAvatar(c echo.Context) error {
	userID := application.GetClaimsUserID(c)
	request := new(struct {
		AvatarData string `json:"avatar" validate:"required"`
	})
	c.Bind(&request)
	err := h.service.UpdateUserAvatar(userID, request.AvatarData)
	if err != nil {
		return errors.Trace(err)
	}
	return c.NoContent(http.StatusOK)
}
