package main

import (
	"context"
	"time"

	"github.com/koioannis/chatter/internal/adapters/events"
	"github.com/koioannis/chatter/internal/adapters/http"
	"github.com/koioannis/chatter/internal/adapters/store"
	"github.com/koioannis/chatter/internal/core/services"
	"github.com/koioannis/chatter/pkg/logging"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	app := fx.New(
		fx.Provide(
			func(logger *logrus.Logger) *echo.Echo {
				e := echo.New()
				e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
					LogURI:    true,
					LogStatus: true,
					LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
						logger.WithFields(logrus.Fields{
							"URI":    values.URI,
							"status": values.Status,
							"method": values.Method,
						}).Info("Request")
						return nil
					},
				}))
				e.Static("/static", "static/dist")

				return e
			},
		),
		fx.Invoke(
			func(lc fx.Lifecycle, e *echo.Echo) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go e.Start(":3000")
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return e.Shutdown(ctx)
					},
				})
			},
		),
		http.Module,
		logging.Module,
		store.Module,
		services.Module,
		events.Module,
		fx.WithLogger(func(logger *logrus.Logger) fxevent.Logger {
			return logging.NewLoggerAdapter(logger)
		}),
	)

	app.Run()

	time.Sleep(time.Second * 2)
}
