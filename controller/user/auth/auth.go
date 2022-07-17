package auth

import (
	"net/http"

	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/labstack/echo/v4"
)

type ControllerUser struct {
	Svc domains.AuthServiceUser
}

// RegisterUser godoc
// @Summary Register User
// @Description Create User Admin
// @Tags Auth User
// @accept json
// @Produce json
// @Router /user/register [post]
// @Param data body entities.RegisterUser true "required"
// @Success 201 {object} entities.User
// @Failure 401 {object} entities.User
// @Failure 500 {object} entities.User
func (co *ControllerUser) RegisterUser(c echo.Context) error {
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

// LoginUser godoc
// @Summary Login User
// @Description Login User
// @Tags Auth User
// @accept json
// @Produce json
// @Router /user/login [post]
// @Param data body entities.LoginUser true "required"
// @Success 200 {object} entities.User
// @Failure 400 {object} entities.User
// @Failure 401 {object} entities.User
// @Failure 500 {object} entities.User
func (co *ControllerUser) LoginUser(c echo.Context) error {
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
