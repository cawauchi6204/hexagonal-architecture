package routers

import (
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/presentation/handlers"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/service_locater"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(s *service_locater.ServiceLocater) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	userHandler := &handlers.UserHandler{
		ServiceLocater: s,
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "一応動いてます!")
	})
	e.GET("/users", userHandler.List)

	return e
}
