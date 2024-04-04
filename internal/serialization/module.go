package serialization

import "go.uber.org/fx"

var Module = fx.Provide(
	fx.Annotate(
		func() JSONMessageSerializerDeserializer {
			return JSONMessageSerializerDeserializer{}
		},
		fx.As(new(MessageSerializer)),
		fx.As(new(MessageDeserializer)),
	),
)
