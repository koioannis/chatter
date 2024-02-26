package handlers

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/app/chat"
	"github.com/koioannis/chatter/internal/web/components"
	"github.com/labstack/echo/v4"
)

type RoomHandler struct{}

func RegisterRoomHandler(e *echo.Echo) {
	h := &RoomHandler{}
	e.GET("/room", h.get)
	e.POST("/room", h.create)
	e.GET("/create-room", h.getCreateRoom)
}

func (h *RoomHandler) create(c echo.Context) error {
	req := struct {
		Name string `form:"name"`
	}{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, err)
	}

	room := chat.NewRoom(uuid.New(), req.Name)
	fmt.Println("Room created: ", room)
	return components.Room(room).Render(c.Request().Context(), c.Response().Writer)

}

func (h *RoomHandler) getCreateRoom(c echo.Context) error {
	return components.AddRoomModal().Render(c.Request().Context(), c.Response().Writer)
}

func (h *RoomHandler) get(c echo.Context) error {
	rooms := []*chat.Room{}
	for i := 0; i < 5; i++ {
		rooms = append(rooms, chat.NewRoom(uuid.New(), fmt.Sprintf("%d Room", i)))
	}
	return components.Rooms(rooms).Render(c.Request().Context(), c.Response().Writer)
}
