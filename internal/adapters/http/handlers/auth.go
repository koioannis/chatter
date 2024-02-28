package handlers

import (
	"github.com/labstack/echo/v4"
)

func DummyAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("username")
		if err != nil {
			return echo.NewHTTPError(401, "Unauthorized")
		}

		c.Set("username", cookie.Value)
		return next(c)
	}
}
