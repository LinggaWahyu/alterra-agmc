package http

import (
	"alterra-agmc-day-5-6/internal/app/auth"
	"alterra-agmc-day-5-6/internal/app/book"
	"alterra-agmc-day-5-6/internal/app/user"
	"alterra-agmc-day-5-6/internal/factory"

	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	v1 := e.Group("/v1")

	book.NewHandler(f).Route(v1.Group("/books"))
	user.NewHandler(f).Route(v1.Group("/users"))
	auth.NewHandler(f).Route(v1)
}
