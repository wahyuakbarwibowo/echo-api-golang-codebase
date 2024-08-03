package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	userID := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"id":   userID,
		"name": "Example User",
	})
}
