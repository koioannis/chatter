package serialization

import (
	"encoding/json"

	"github.com/koioannis/chatter/internal/core/domain"
)

type MessageSerializer interface {
	Serialize(message *domain.Message) ([]byte, error)
}

type MessageDeserializer interface {
	Deserialize(bb []byte) (*domain.Message, error)
}

type JSONMessageSerializerDeserializer struct{}

func (s JSONMessageSerializerDeserializer) Serialize(message *domain.Message) ([]byte, error) {
	return json.Marshal(message)
}

func (s JSONMessageSerializerDeserializer) Deserialize(bb []byte) (*domain.Message, error) {
	var m *domain.Message
	err := json.Unmarshal(bb, &m)

	return m, err
}
