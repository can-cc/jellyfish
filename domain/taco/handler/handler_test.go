package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fwchen/jellyfish/application/middleware"
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/repository/mock"
	"github.com/fwchen/jellyfish/domain/taco/service"
	"github.com/fwchen/jellyfish/util"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandler_GetTacos(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tacos", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", &jwt.Token{Claims: &middleware.JwtAppClaims{
		SignData: middleware.SignData{ID: "u123"},
	}})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	mockRepo.EXPECT().ListTacos("u123", gomock.Any()).Return([]taco.Taco{
		{
			ID:        "id_123",
			CreatorID: "u123",
			Content:   "watch tv",
			Detail:    "watch AC",
			Status:    "Doing",
			Type:      "Task",
			Deadline:  nil,
			CreatedAt: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
			UpdateAt:  nil,
		},
	}, nil)

	h := &handler{service.NewTacoApplicationService(mockRepo)}

	if assert.NoError(t, h.GetTacos(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `[{"id":"id_123","creatorID":"u123","content":"watch tv","detail":"watch AC","status":"Doing","type":"Task","deadline":null,"createdAt":"2020-01-02T00:00:00Z","updatedAt":null}]`, rec.Body.String())
	}
}

func TestHandler_CreateTaco(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/tacos", strings.NewReader(`{"content": "Read book", "detail": "《Refactor》", "deadline": "2020-01-16T00:00:00Z"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", &jwt.Token{Claims: &middleware.JwtAppClaims{
		SignData: middleware.SignData{ID: "u123"},
	}})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	mockRepo.EXPECT().InsertTaco(&taco.Taco{
		CreatorID: "u123",
		Content:   "Read book",
		Detail:    "《Refactor》",
		Status:    "Doing",
		Type:      "Task",
		Deadline:  util.PointerTime(time.Date(2020, 1, 16, 0, 0, 0, 0, time.UTC)),
	}).Return(util.PointerStr("tid_1"), nil)

	h := &handler{service.NewTacoApplicationService(mockRepo)}

	if assert.NoError(t, h.CreateTaco(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
