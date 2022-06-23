package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	Svc domains.AuthService
}

// Register godoc
// @Summary Register User Admin
// @Description Create User Admin
// @Tags Auth
// @accept json
// @Produce json
// @Router /admin/register [post]
// @Param data body entities.RegisterAdmin true "required"
// @Success 201 {object} entities.Admin
// @Failure 401 {object} entities.Admin
// @Failure 500 {object} entities.Admin
func (co *AuthController) Register(c echo.Context) error {
	admins := entities.Admin{}
	err := c.Bind(&admins)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	er := co.Svc.RegisterService(admins)

	if er != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "user UnRegistered",
			"data":    er.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success register admin",
		"data":    admins,
	})
}

// Login godoc
// @Summary Login User Admin
// @Description Login User Admin
// @Tags Auth
// @accept json
// @Produce json
// @Router /admin/login [post]
// @Param data body entities.LoginAdmin true "required"
// @Success 200 {object} entities.Admin
// @Failure 400 {object} entities.Admin
// @Failure 401 {object} entities.Admin
// @Failure 500 {object} entities.Admin
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

// GetUser godoc
// @Summary Get User With name
// @Description Admin can get the user name
// @Tags Auth
// @accept json
// @Produce json
// @Router /admin/user/{name} [get]
// @Param name path string true "name"
// @Success 200 {object} entities.Admin
// @Failure 404 {object} entities.Admin
// @Security JWT
func (co *AuthController) GetUser(c echo.Context) error {
	name := c.Param("name")

	admins, er := co.Svc.GetUserService(name)

	if er != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"data":    er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user admin by name",
		"data":    admins,
	})
}
