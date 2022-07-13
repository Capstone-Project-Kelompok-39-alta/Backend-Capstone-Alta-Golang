package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
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
		"message":    "success login admin",
		"id_pegawai": adminLogin.IdPegawai,
		"data":       token,
	})
}

// GetUser godoc
// @Summary Get User With name
// @Description Admin can get by ID
// @Tags Auth
// @accept json
// @Produce json
// @Router /admin/user/{id_pegawai} [get]
// @param id_pegawai path int true "id_pegawai"
// @Success 200 {object} entities.Admin
// @Failure 404 {object} entities.Admin
// @Security JWT
func (co *AuthController) GetUser(c echo.Context) error {
	id_pegawai := c.Param("id_pegawai")
	intID, _ := strconv.Atoi(id_pegawai)

	admins, er := co.Svc.GetUserService(intID)

	if er != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"data":    er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user admin by id_pegawai",
		"data":    admins,
	})
}

// UpdateUser godoc
// @Summary Update User Admin
// @Description Update User Admin
// @Tags Auth
// @accept json
// @Produce json
// @Router /admin/user/{id_pegawai} [put]
// @param id_pegawai path int true "id_pegawai"
// @param data body entities.Admin true "required"
// @Success 200 {object} entities.Admin
// @Failure 500 {object} entities.Admin
// @Security JWT
func (co *AuthController) UpdateUser(c echo.Context) error {
	var admin entities.Admin
	id_pegawai, _ := strconv.Atoi(c.Param("id_pegawai"))

	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error to update the admin",
		})
	}

	auth, er := co.Svc.UpdateUserService(id_pegawai)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error to update the admin",
			"data":    er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success edited the admin",
		"data":    auth,
	})
}
