package auth

import (
	"context"

	"github.com/labstack/echo/v4"
)

type key string

var usernameKey = key("username")

func GetCurrentUser(c context.Context) string {
	val := c.Value(usernameKey)
	if val == nil {
		return ""
	}

	return val.(string)
}

func DummyAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("username")
		if err != nil {
			c.Redirect(302, "/login")
			return nil
		}

		c.Set("username", cookie.Value)

		c.SetRequest(
			c.Request().WithContext(
				context.WithValue(c.Request().Context(), usernameKey, cookie.Value),
			),
		)
		return next(c)
	}
}
