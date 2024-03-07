package handlers

import (
	"github.com/a-h/templ"
	"github.com/koioannis/chatter/internal/adapters/http/templates"
	"github.com/labstack/echo/v4"
)

func render(component templ.Component, ctx echo.Context) error {
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)

}

func renderWithIndex(component templ.Component, ctx echo.Context) error {
	return templates.Index().Render(templ.WithChildren(ctx.Request().Context(), component), ctx.Response().Writer)
}
