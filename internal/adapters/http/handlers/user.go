package handlers

import (
	"net/http"

	"github.com/koioannis/chatter/internal/adapters/http/templates"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func RegisterUserHandler(e *echo.Echo) {
	h := &UserHandler{}
	e.POST("/login", h.login)
}

func (h *UserHandler) login(c echo.Context) error {
	req := struct {
		Username string `form:"username"`
	}{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, err)
	}

	cookie := &http.Cookie{
		Name:  "username",
		Value: req.Username,
	}

	c.SetCookie(cookie)
	return templates.Home().Render(c.Request().Context(), c.Response().Writer)
}
