package admin

import (
	"errors"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/testing/service_test/auth"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterController(t *testing.T) {
	svc := auth.AuthService{}

	AuthController := AuthController{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		svc.On("RegisterService", mock.Anything).Return(nil).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/admin/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := AuthController.Register(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		svc.On("RegisterService", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/admin/register", io.Reader(strings.NewReader(`{"Status" : "Internal Server Error"}`)))
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := AuthController.Register(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("StatusUnauthorized", func(t *testing.T) {
		svc.On("RegisterService", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/admin/register", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := AuthController.Register(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 401, w.Result().StatusCode)
	})
}

func TestLoginController(t *testing.T) {
	svc := auth.AuthService{}
	adminLogin := new(auth.AuthService)
	adminLoginData := entities.Admin{
		Name:      "Name Testing",
		IdPegawai: 12345678,
		Email:     "emailTesting@gmail.com",
		Password:  "Password Testing",
	}

	AuthControllers := AuthController{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		adminLogin.On("LoginService", mock.Anything, mock.Anything).Return("success login", http.StatusOK).Once()
		adminLogin.LoginService(adminLoginData.IdPegawai, adminLoginData.Password)
		svc.On("LoginService", mock.Anything, mock.Anything).Return("success login", http.StatusOK).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/admin/login", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := AuthControllers.Login(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("StatusBadRequest", func(t *testing.T) {
		adminLogin.On("LoginService", mock.Anything, mock.Anything).Return("StatusBadRequest", http.StatusBadRequest).Once()
		adminLogin.LoginService(adminLoginData.IdPegawai, adminLoginData.Password)
		svc.On("LoginService", mock.Anything, mock.Anything).Return("StatusBadRequest", http.StatusBadRequest).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/admin/login", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := AuthControllers.Login(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 400, w.Result().StatusCode)
	})
}

func TestGetUser(t *testing.T) {
	svc := auth.AuthService{}
	authService := new(auth.AuthService)
	adminData := entities.Admin{
		Name:      "Name Testing",
		IdPegawai: 12345678,
		Email:     "emailTesting@gmail.com",
		Password:  "Password Testing",
	}

	AuthController := AuthController{
		Svc: &svc,
	}

	t.Run("StatusOK", func(t *testing.T) {
		authService.On("GetUserService", mock.Anything).Return(adminData, nil).Once()
		authService, err := authService.GetUserService(int(adminData.IdPegawai))
		svc.On("GetUserService", mock.Anything).Return(adminData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/admin/user/:id_pegawai", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := AuthController.GetUser(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, authService, adminData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("StatusNotFound", func(t *testing.T) {
		authService.On("GetUserService", mock.Anything).Return(adminData, errors.New("error to get admin by id_pegawai")).Once()
		authService, err := authService.GetUserService(int(adminData.IdPegawai))
		svc.On("GetUserService", mock.Anything).Return(adminData, errors.New("error to get admin by id_pegawai")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/admin/user/:id_pegawai", io.Reader(strings.NewReader(`{"Status" : "Not Found"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := AuthController.GetUser(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, authService, adminData)
		assert.Equal(t, 404, w.Result().StatusCode)
		assert.Error(t, err)
	})
}
