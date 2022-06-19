package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/admin"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	Svc domains.AuthService
}

func (co *AuthController) Register(c echo.Context) error {
	admin := entities.Admin{}
	err := c.Bind(&admin)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	er := co.Svc.RegisterService(admin)

	if er != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "user UnRegistered",
			"data":    er.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success register admin",
		"data":    admin,
	})
}

func (co *AuthController) Login(c echo.Context) error {
	adminLogin := entities.Admin{}

	err := c.Bind(&adminLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "bad request",
			"error":  err.Error(),
		})
	}

	token, statusCode := co.Svc.LoginService(adminLogin.IdPegawai, adminLogin.Password)

	if statusCode == http.StatusUnauthorized {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "your id_pegawai and password wrong",
			"statusCode": statusCode,
		})
	} else if statusCode == http.StatusInternalServerError {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":    "internal server error",
			"statusCode": statusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login admin",
		"data":    token,
	})
}

func (co *AuthController) GetUser(c echo.Context) error {
	name := c.Param("name")

	admin, er := co.Svc.GetUserService(name)

	if name != admin.Name {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user name not found in database",
			"status":  http.StatusNotFound,
		})
	}

	if er != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"data":    er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user admin by name",
		"data":    admin,
	})
}
