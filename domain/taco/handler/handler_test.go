package handler

import (
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/repository"
	"github.com/fwchen/jellyfish/domain/taco/repository/mock"
	"github.com/fwchen/jellyfish/domain/taco/service"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_GetTacos(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tacos", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("App", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InUxIiwiZXhwIjoxNTgxMTczNzE2fQ.9SFPKfIRKyWwfPFBxlk1YVGJnL8l17BRj_ZkRIawQaA")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	mockRepo.EXPECT().ListTacos("u123", repository.ListTacoFilter{Statues: []taco.Status{taco.Doing}}).Return([]taco.Taco{
		{
			ID:        "id_123",
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
		assert.Equal(t, `{"id":"123-456","username":"oyx"}`, rec.Body.String())
	}
}
