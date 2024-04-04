package events

import (
	"context"

	ports "github.com/koioannis/chatter/internal/ports/events"
	"github.com/koioannis/chatter/internal/serialization"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Module("events",
	fx.Provide(
		func(lc fx.Lifecycle, serializer serialization.MessageSerializer) ports.MessagePublisher {
			publisher := NewKafkaMessagePublisher(serializer)
			lc.Append(
				fx.Hook{
					OnStop: func(ctx context.Context) error {
						return publisher.Stop()
					},
				},
			)

			return publisher
		},
		func(lc fx.Lifecycle, deserializer serialization.MessageDeserializer, logger *logrus.Logger) ports.MessageSubscriber {
			subscriber := NewKafkaMessageSubscriber(deserializer, logger)

			subscriberCtx, cancel := context.WithCancel(context.Background())
			lc.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							if err := subscriber.Start(subscriberCtx); err != nil && err != context.Canceled {
								log.Error("Subscriber stopped with error: ", err)
							}
						}()

						return nil
					},
					OnStop: func(ctx context.Context) error {
						logger.Info("Stopping subscriber")
						cancel()
						return nil
					},
				},
			)
			return subscriber

		},
	),
)
