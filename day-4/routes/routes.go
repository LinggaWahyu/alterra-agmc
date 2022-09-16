package routes

import (
	"alterra-agmc-day-4/constants"
	"alterra-agmc-day-4/controllers"
	m "alterra-agmc-day-4/middlewares"
	"alterra-agmc-day-4/util"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Validator = &util.CustomValidator{Validator: validator.New()}

	m.LogMiddlewares(e)

	publicRoutes := e.Group("/v1")
	{
		publicRoutes.POST("/login", controllers.LoginUserControllers)

		publicRoutes.POST("/users", controllers.CreateUserControllers)

		publicRoutes.GET("/books", controllers.GetBookControllers)
		publicRoutes.GET("/books/:id", controllers.GetBookByIdControllers)
	}

	authenticatedRoutes := e.Group("/v1")
	{
		authenticatedRoutes.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

		authenticatedRoutes.GET("/users", controllers.GetUserControllers)
		authenticatedRoutes.GET("/users/:id", controllers.GetUserByIdControllers)
		authenticatedRoutes.PUT("/users/:id", controllers.UpdateUserControllers)
		authenticatedRoutes.DELETE("/users/:id", controllers.DeleteUserControllers)

		authenticatedRoutes.POST("/books", controllers.CreateBookControllers)
		authenticatedRoutes.PUT("/books/:id", controllers.UpdateBookControllers)
		authenticatedRoutes.DELETE("/books/:id", controllers.DeleteBookControllers)
	}

	return e
}
