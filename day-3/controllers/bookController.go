package controllers

import (
	"alterra-agmc-day-3/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBookControllers(c echo.Context) error {
	books := []models.Book{
		{
			ID:        1,
			Title:     "Laskar Pelangi",
			Author:    "Andrea Hirata",
			Publisher: "Bentang Pustaka (Yogyakarta)",
		},
		{
			ID:        2,
			Title:     "Sang Pemimpi",
			Author:    "Andrea Hirata",
			Publisher: "Yogyakarta: Bentang Pustaka",
		},
		{
			ID:        3,
			Title:     "Edensor",
			Author:    "Andrea Hirata",
			Publisher: "Yogyakarta: Bentang Pustaka",
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func GetBookByIdControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book := models.Book{
		ID:        uint(id),
		Title:     "Perahu Kertas",
		Author:    "Dee",
		Publisher: "Bentang Pustaka (Yogyakarta)",
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func CreateBookControllers(c echo.Context) error {
	var book models.Book
	err := c.Bind(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func UpdateBookControllers(c echo.Context) error {
	book := models.Book{
		Title:     "Perahu Kertas",
		Author:    "Dee",
		Publisher: "Bentang Pustaka (Yogyakarta)",
	}

	id, _ := strconv.Atoi(c.Param("id"))
	book.ID = uint(id)

	err := c.Bind(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func DeleteBookControllers(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "book with ID " + id + " has been deleted",
	})
}
