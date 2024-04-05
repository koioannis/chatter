package auth

import (
	"context"

	"github.com/labstack/echo/v4"
)

type key string

var usernameKey = key("username")

func NewContext(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, usernameKey, username)
}

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

		c.SetRequest(
			c.Request().WithContext(
				NewContext(c.Request().Context(), cookie.Value),
			),
		)
		return next(c)
	}
}
