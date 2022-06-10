package auth

import (
	"net/http"

	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/user"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Svc domains.AuthService
}

func (co *AuthController) Register(c echo.Context) error {
	user := entities.User{}
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	er := co.Svc.RegisterService(user)

	if er != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "user UnRegistered",
			"data":    er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success register user",
		"data":    user,
	})
}

func (co *AuthController) Login(c echo.Context) error {
	UserLogin := entities.User{}

	err := c.Bind(&UserLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "bad request",
			"error":  err.Error(),
		})
	}

	token, statusCode := co.Svc.LoginService(UserLogin.Email, UserLogin.Password)

	if statusCode == http.StatusUnauthorized {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "your email and password wrong",
			"statusCode": statusCode,
		})
	} else if statusCode == http.StatusInternalServerError {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":    "internal server error",
			"statusCode": statusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login user",
		"data":    token,
	})
}
