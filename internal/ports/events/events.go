package events

import (
	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/core/domain"
)

type MessagePublisher interface {
	Publish(*domain.Message) error
}

type MessageSubscriber interface {
	Subscribe(roomID uuid.UUID) <-chan *domain.Message
}
