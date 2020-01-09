package handler

import (
	"github.com/dchest/captcha"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/fwchen/jellyfish/domain/visitor"
	"github.com/fwchen/jellyfish/domain/visitor/repository/mock"
	visitorUtil "github.com/fwchen/jellyfish/domain/visitor/util"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Login(t *testing.T) {
	captchaStore := captcha.NewMemoryStore(captcha.CollectNum, captcha.Expiration)
	captchaStore.Set("captcha_1", []byte{5, 5, 0, 5, 5, 5})
	captcha.SetCustomStore(captchaStore)

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		strings.NewReader(`{"username": "moon", "password": "cloud", "captcha": "550555", "captchaID": "captcha_1"}`),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	hash, _ := visitorUtil.GenerateFromPassword("cloud2")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindUserIDAndHash("moon").Return("id1", hash, nil)

	h := NewHandler(mockRepo, &configs.ApplicationConfig{
		JwtHeaderKey: "Test-Authorization",
	})

	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, rec.Header().Get("Test-Authorization"))
	}
}

func TestHandler_SignUp(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/signup",
		strings.NewReader(`{"username": "moon", "password": "cloud"}`),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	mockRepo.EXPECT().Save(&visitor.Visitor{
		Name:        "moon",
		Password:    "cloud",
		IsCertified: true,
	}, gomock.Any()).Return("id_123", nil)

	h := NewHandler(mockRepo, &configs.ApplicationConfig{})

	if assert.NoError(t, h.SignUp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestHandler_GenCaptcha(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/captcha",
		nil,
	)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := NewHandler(nil, &configs.ApplicationConfig{})

	if assert.NoError(t, h.GenCaptcha(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotNil(t, rec.Body.String())
	}
}
