package events

import (
	ports "github.com/koioannis/chatter/internal/ports/events"
	"go.uber.org/fx"
)

var Module = fx.Module("events",
	fx.Provide(
		fx.Annotate(
			NewInmemoryPublisherSubscriber,
			fx.As(new(ports.MessagePublisher)),
			fx.As(new(ports.MessageSubscriber)),
		),
	),
)
