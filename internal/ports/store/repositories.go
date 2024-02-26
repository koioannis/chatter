package store

import "github.com/koioannis/chatter/internal/core/domain"

type RoomRepository interface {
	Get(name string) (*domain.Room, error)
	GetAll() ([]*domain.Room, error)
	Create(room *domain.Room) error
}
