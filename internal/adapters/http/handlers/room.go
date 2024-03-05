package handlers

import (
	"github.com/koioannis/chatter/internal/adapters/http/auth"
	"github.com/koioannis/chatter/internal/adapters/http/templates"
	"github.com/koioannis/chatter/internal/core/domain"
	"github.com/koioannis/chatter/internal/core/services"
	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	service *services.RoomService
}

func RegisterRoomHandler(e *echo.Echo, service *services.RoomService) {
	g := e.Group("/room", auth.DummyAuth)

	h := &RoomHandler{
		service: service,
	}
	g.GET("", h.get)
	g.POST("", h.create)
	g.GET("/create", h.getCreateRoom)
}

func (h *RoomHandler) create(c echo.Context) error {
	req := struct {
		Name string `form:"name"`
	}{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, err)
	}

	room, err := h.service.Create(req.Name)
	if err != nil {
		c.Response().Header().Add("HX-Reswap", "outerHTML")
		switch err {
		case domain.ErrRoomAlreadyExists:
			c.Response().WriteHeader(409)
		case domain.ErrInvalidRoomName:
			c.Response().WriteHeader(422)
		default:
			c.Response().WriteHeader(500)
		}

		return templates.AddRoomModal(req.Name, err).Render(c.Request().Context(), c.Response().Writer)
	}

	return templates.Room(room).Render(c.Request().Context(), c.Response().Writer)

}

func (h *RoomHandler) getCreateRoom(c echo.Context) error {
	return templates.AddRoomModal("", nil).Render(c.Request().Context(), c.Response().Writer)
}

func (h *RoomHandler) get(c echo.Context) error {
	rooms := h.service.GetAll()
	return templates.Rooms(rooms).Render(c.Request().Context(), c.Response().Writer)
}
