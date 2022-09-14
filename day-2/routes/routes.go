package routes

import (
	"alterra-agmc-day-2/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/v1")

	v1.GET("/users", controllers.GetUserControllers)
	v1.GET("/users/:id", controllers.GetUserByIdControllers)
	v1.POST("/users", controllers.CreateUserControllers)
	v1.PUT("/users/:id", controllers.UpdateUserControllers)
	v1.DELETE("/users/:id", controllers.DeleteUserControllers)

	v1.GET("/books", controllers.GetBookControllers)
	v1.GET("/books/:id", controllers.GetBookByIdControllers)
	v1.POST("/books", controllers.CreateBookControllers)
	v1.PUT("/books/:id", controllers.UpdateBookControllers)
	v1.DELETE("/books/:id", controllers.DeleteBookControllers)

	return e
}
