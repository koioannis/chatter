package store

import (
	ports "github.com/koioannis/chatter/internal/ports/store"
	"go.uber.org/fx"
)

var Module = fx.Module("persistance",
	fx.Provide(
		fx.Annotate(
			NewInMemoryRoomRepository,
			fx.As(new(ports.RoomRepository)),
		),
	),
)
