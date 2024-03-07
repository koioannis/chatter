package handlers

import (
	"net/http"

	"github.com/koioannis/chatter/internal/adapters/http/templates"
	"github.com/labstack/echo/v4"
)

var cookieName = "username"

type UserHandler struct{}

func RegisterUserHandler(e *echo.Echo) {
	h := &UserHandler{}
	e.POST("/login", h.login)
	e.GET("/login", h.getLogin)
}

func (h *UserHandler) getLogin(c echo.Context) error {

	_, err := c.Cookie(cookieName)
	if err == nil {
		c.Response().Header().Set("Location", "/home")
		c.Response().WriteHeader(302)
		return renderWithIndex(templates.Home(), c)
	}

	return renderWithIndex(templates.Login(), c)
}

func (h *UserHandler) login(c echo.Context) error {

	req := struct {
		Username string `form:"username"`
	}{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, err)
	}

	cookie := &http.Cookie{
		Name:  cookieName,
		Value: req.Username,
	}

	c.SetCookie(cookie)
	return templates.Home().Render(c.Request().Context(), c.Response().Writer)
}
