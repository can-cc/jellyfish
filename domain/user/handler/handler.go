package handler

import (
	"github.com/juju/errors"
	"github.com/labstack/echo"
	"jellyfish/application/middleware"
	"jellyfish/domain/user/repository"
	userService "jellyfish/domain/user/service"
	"jellyfish/service"
	"net/http"
)

type handler struct {
	service *userService.ApplicationService
}

func NewHandler(userRepo repository.Repository, imageStorageService *service.ImageStorageService) *handler {
	return &handler{service: userService.NewApplicationService(userRepo, imageStorageService)}
}

func (h *handler) GetUserInfo(c echo.Context) error {
	userID := middleware.GetClaimsUserID(c)
	userInfo, err := h.service.GetUserInfo(userID)
	if err != nil {
		return errors.Trace(err)
	}
	return c.JSON(http.StatusOK, userInfo)
}

func (h *handler) UpdateUserAvatar(c echo.Context) error {
	userID := middleware.GetClaimsUserID(c)
	request := new(struct {
		Avatar string `json:"avatar" validate:"required"`
	})
	err := c.Bind(&request)
	if err != nil {
		return errors.Trace(err)
	}
	err = h.service.UpdateUserAvatar(userID, request.Avatar)
	if err != nil {
		return errors.Trace(err)
	}
	return c.NoContent(http.StatusOK)
}
