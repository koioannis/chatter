package ws

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/adapters/http/auth"
	"github.com/koioannis/chatter/internal/adapters/http/templates"
	"github.com/koioannis/chatter/internal/core/domain"
	"github.com/koioannis/chatter/internal/ports/events"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type WsHandler struct {
	subscriber events.MessageSubscriber
}

func RegisterWsHandler(e *echo.Echo, subscriber events.MessageSubscriber) {
	handler := &WsHandler{
		subscriber: subscriber,
	}

	e.GET("/messages/ws/:roomID", handler.OnNewConn, auth.DummyAuth)
}

func (h *WsHandler) OnNewConn(c echo.Context) error {
	roomID := c.Param("roomID")
	roomUUID := uuid.MustParse(roomID)
	username := auth.GetCurrentUser(c.Request().Context())

	websocket.Handler(func(conn *websocket.Conn) {
		ctx, cancel := context.WithCancel(context.Background())
		defer conn.Close()
		go h.monitorConnection(cancel, conn)

		go func() {
			for {
				select {
				case msg := <-h.subscriber.Subscribe(roomUUID):
					if msg.Sender == username {
						continue
					}

					if err := templates.Messages([]*domain.Message{msg}, true).Render(context.Background(), conn); err != nil {
						return
					}
				case <-ctx.Done():
					fmt.Println("closing write!")
					return
				}
			}
		}()

		<-ctx.Done()
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func (h *WsHandler) monitorConnection(cancel context.CancelFunc, conn *websocket.Conn) {
	for {
		var message []byte
		_, err := conn.Read(message)
		if err != nil {
			cancel()
			return
		}
	}
}
