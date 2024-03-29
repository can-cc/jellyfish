package handler

import (
	"github.com/dgrijalva/jwt-go"
	"jellyfish/application/middleware"
	"jellyfish/domain/user"
	"jellyfish/domain/user/repository/mock"
	userService "jellyfish/domain/user/service"
	"jellyfish/util"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_GetUserInfo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/u123", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", &jwt.Token{Claims: &middleware.JwtAppClaims{
		SignData: middleware.SignData{ID: "u123"},
	}})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID("u123").Return(&user.AppUser{
		ID:        util.PointerStr("123-456"),
		Username:  util.PointerStr("oyx"),
		CreatedAt: func(t time.Time) *time.Time { return &t }(time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC)),
		UpdatedAt: nil,
	}, nil)

	h := &handler{userService.NewApplicationService(mockRepo, nil)}

	if assert.NoError(t, h.GetUserInfo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"id":"123-456","username":"oyx","avatar":""}
`, rec.Body.String())
	}
}
