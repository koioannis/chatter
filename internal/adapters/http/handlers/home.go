package handlers

import (
	"github.com/koioannis/chatter/internal/adapters/http/auth"
	"github.com/koioannis/chatter/internal/adapters/http/templates"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func RegisterHomeHandler(e *echo.Echo) {
	g := e.Group("/home", auth.DummyAuth)

	h := &HomeHandler{}
	g.GET("", h.get)
}

func (h *HomeHandler) get(c echo.Context) error {
	return renderWithIndex(templates.Home(), c)
}
