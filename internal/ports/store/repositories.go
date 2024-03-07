package store

import (
	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/core/domain"
)

type RoomRepository interface {
	Get(name string) (*domain.Room, error)
	GetById(id uuid.UUID) (*domain.Room, error)
	GetAll() ([]*domain.Room, error)
	Create(room *domain.Room) error
}

type MessageRepository interface {
	Create(message *domain.Message) error
	GetAllByRoomId(roomId uuid.UUID) ([]*domain.Message, error)
}
