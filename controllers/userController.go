package controllers

import (
	"net/http"
	"test/lib/database"
	"test/models"
	"test/responses"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	inputUser := models.User{}
	c.Bind(&inputUser)

	inputUser.Role = c.Param("role")

	user, err := database.RegisterUser(inputUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "failed to register new account"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse("success", "success to register new account", user))
}

func LoginUserController(c echo.Context) error {
	inputData := models.LoginUser{}
	c.Bind(&inputData)

	loggedUser, rowAffected, err := database.LoginUser(inputData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "failed to login"))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, responses.FailedResponse("success", "user data not found"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse("success", "success to login", loggedUser))
}
