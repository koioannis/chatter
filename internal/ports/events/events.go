package events

import (
	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/core/domain"
)

type MessagePublisher interface {
	Publish(message *domain.Message) error
}

type MessageSubscriber interface {
	Subscribe(roomID uuid.UUID, clientID string) <-chan *domain.Message
	Unsubscribe(roomID uuid.UUID, clientID string)
}
