package app

import (
	"echo-api-golang-codebase/internal/infrastructure/http"

	"github.com/labstack/echo/v4"
	"honnef.co/go/tools/config"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(cfg *config.Config) *Server {
	e := echo.New()
	http.RegisterRoutes(e)
	return &Server{echo: e}
}

func (s *Server) Run() error {
	return s.echo.Start(":8080")
}
