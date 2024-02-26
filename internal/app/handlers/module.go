package handlers

import "go.uber.org/fx"

var Module = fx.Module("handlers",
	fx.Invoke(RegisterIndexHandler),
	fx.Invoke(RegisterUserHandler),
	fx.Invoke(RegisterRoomHandler),
)
