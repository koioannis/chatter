package http

import (
	"github.com/koioannis/chatter/internal/adapters/http/handlers"
	"github.com/koioannis/chatter/internal/adapters/http/ws"
	"go.uber.org/fx"
)

var Module = fx.Module("handlers",
	fx.Invoke(handlers.RegisterIndexHandler),
	fx.Invoke(handlers.RegisterUserHandler),
	fx.Invoke(handlers.RegisterRoomHandler),
	fx.Invoke(handlers.RegisterHomeHandler),
	fx.Invoke(handlers.RegisterMessageHandler),
	fx.Invoke(ws.RegisterWsHandler),
)
