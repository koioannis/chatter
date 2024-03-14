package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/adapters/http"
	"github.com/koioannis/chatter/internal/adapters/http/auth"
	"github.com/koioannis/chatter/internal/adapters/http/templates"
	"github.com/koioannis/chatter/internal/adapters/store"
	"github.com/koioannis/chatter/internal/core/domain"
	"github.com/koioannis/chatter/internal/core/services"
	"github.com/koioannis/chatter/pkg/logging"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"golang.org/x/net/websocket"
)

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		fmt.Println("New connection!")
		closeChan := make(chan struct{})
		defer ws.Close()
		ticker := time.NewTicker(time.Second * 5)
		wg := &sync.WaitGroup{}

		wg.Add(2)
		go func(wg *sync.WaitGroup) {
			var message []byte
			_, err := ws.Read(message)

			// If an error occurred, check if it's because the client closed the connection
			if err != nil {
				fmt.Println("Client has closed the connection or an error occurred:", err)
				wg.Done()
				closeChan <- struct{}{}
				return // Exit the for loop and close the connection
			}

		}(wg)
		go func(wg *sync.WaitGroup) {
			for {
				select {
				case <-ticker.C:
					fmt.Println("Writing")
					m, err := domain.NewMessage(uuid.New(), "John", uuid.New(), "hi!")
					if err != nil {
						log.Fatal(err)
					}
					if err := templates.Messages([]*domain.Message{m}, true).Render(context.Background(), ws); err != nil {
						fmt.Println("Error while writing to client:", err)
						return // Exit the for loop and close the connection
					}
				case <-closeChan:
					fmt.Println("closing write!")
					wg.Done()
					return
				}

			}
		}(wg)

		wg.Wait()
		fmt.Println("handler done")
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

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
						e.GET("/ws", hello, auth.DummyAuth)
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
		fx.WithLogger(func(logger *logrus.Logger) fxevent.Logger {
			return logging.NewLoggerAdapter(logger)
		}),
	)

	app.Run()

	time.Sleep(time.Second * 2)
}
