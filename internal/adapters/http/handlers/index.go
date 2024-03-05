package handlers

import (
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

	if isLoggedIn {
		c.Redirect(302, "/home")
	} else {
		c.Redirect(302, "/login")
	}

	return nil
}
