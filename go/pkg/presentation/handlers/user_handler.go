package handlers

import (
	"github.com/labstack/echo/v4"
)

type UserHandler BaseHandler

func (h *UserHandler) List(c echo.Context) error {
	u := h.ServiceLocater.UserUsecase
	users, err := u.List(c.Request().Context())
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, users)
}
