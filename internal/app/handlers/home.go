package handlers

import (
	"github.com/koioannis/chatter/internal/web/components"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func RegisterHomeHandler(e *echo.Echo) {
	h := &HomeHandler{}
	e.GET("/", h.get)
}

func (h *HomeHandler) get(c echo.Context) error {
	_, err := c.Cookie("username")
	isLoggedIn := true
	if err != nil {
		isLoggedIn = false
	}

	return components.Index(isLoggedIn).Render(c.Request().Context(), c.Response().Writer)
}
