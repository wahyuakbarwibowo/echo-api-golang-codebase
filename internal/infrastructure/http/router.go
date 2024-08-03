package http

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	userHandler := handler.NewUserHandler()
	e.GET("/users/:id", userHandler.GetUser)
}
