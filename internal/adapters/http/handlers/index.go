package handlers

import (
	"github.com/koioannis/chatter/internal/adapters/http/templates"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

func RegisterIndexHandler(e *echo.Echo) {
	h := &IndexHandler{}
	e.GET("/", h.get)
}

func (h *IndexHandler) get(c echo.Context) error {
	_, err := c.Cookie("username")
	isLoggedIn := true
	if err != nil {
		isLoggedIn = false
	}

	return templates.Index(isLoggedIn).Render(c.Request().Context(), c.Response().Writer)
}
