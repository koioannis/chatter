package handlers

import (
	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/adapters/http/auth"
	"github.com/koioannis/chatter/internal/adapters/http/templates"
	"github.com/koioannis/chatter/internal/core/services"
	"github.com/koioannis/chatter/internal/dto"
	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	messageService *services.MessageService
	roomService    *services.RoomService
}

func RegisterMessageHandler(messageService *services.MessageService, roomService *services.RoomService, e *echo.Echo) {
	g := e.Group("/room/:room_id/message", auth.DummyAuth)

	h := &MessageHandler{
		messageService: messageService,
		roomService:    roomService,
	}
	g.GET("", h.get)
	g.POST("", h.create)
}

func (h *MessageHandler) get(c echo.Context) error {
	roomIdStr := c.Param("room_id")
	roomId, err := uuid.Parse(roomIdStr)
	if err != nil {
		return err
	}

	room, err := h.roomService.GetById(roomId)
	if err != nil {
		return err
	}

	messages, err := h.messageService.GetAllByRoomId(roomId)
	if err != nil {
		return err
	}

	return render(templates.Chat(room, messages), c)
}

func (h *MessageHandler) create(c echo.Context) error {
	req := struct {
		Message string `form:"message"`
	}{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, err)
	}

	username := auth.GetCurrentUser(c.Request().Context())
	roomID, err := uuid.Parse(c.Param("room_id"))
	if err != nil {
		return echo.NewHTTPError(422, "invalid uuid")
	}

	dto := dto.NewCreateMessageDTO(req.Message, username, roomID)

	msg, err := h.messageService.Create(dto)
	if err != nil {
		return err
	}

	return render(templates.Message(msg), c)
}
