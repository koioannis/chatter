package handlers

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/adapters/http/auth"
	"github.com/koioannis/chatter/internal/adapters/http/templates"
	"github.com/koioannis/chatter/internal/core/services"
	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	s *services.MessageService
}

func RegisterMessageHandler(s *services.MessageService, e *echo.Echo) {
	g := e.Group("/room/:room_id/messages", auth.DummyAuth)

	h := &MessageHandler{s: s}
	g.GET("", h.get)
	g.POST("", h.create)
}

func (h *MessageHandler) get(c echo.Context) error {
	roomIdStr := c.Param("room_id")
	roomId, err := uuid.Parse(roomIdStr)
	if err != nil {
		return err
	}

	messages, err := h.s.GetAllByRoomId(roomId)
	if err != nil {
		return err
	}

	fmt.Println(messages)

	return render(templates.Chat(), c)
}

func (h *MessageHandler) create(c echo.Context) error {
	return nil
}
