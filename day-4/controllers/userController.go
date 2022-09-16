package controllers

import (
	"alterra-agmc-day-4/lib/database"
	"alterra-agmc-day-4/middlewares"
	"alterra-agmc-day-4/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserControllers(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserByIdControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func CreateUserControllers(c echo.Context) error {
	var user models.User

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = c.Validate(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	result, err := database.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   result,
	})
}

func UpdateUserControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	loggedUserId := middlewares.ExtractTokenUserId(c)
	if loggedUserId != int(user.ID) {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": "You don't have permission",
		})
	}

	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = database.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func DeleteUserControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	loggedUserId := middlewares.ExtractTokenUserId(c)
	if loggedUserId != int(user.ID) {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": "You don't have permission",
		})
	}

	err = database.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func LoginUserControllers(c echo.Context) error {
	var user models.User

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = c.Validate(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	user, err = database.LoginUser(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := middlewares.CreateToken(int(user.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"token":  token,
		"user":   user,
	})
}
