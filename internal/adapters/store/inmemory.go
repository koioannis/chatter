package store

import (
	"errors"

	"github.com/koioannis/chatter/internal/core/domain"
)

type InMemoryRoomRepository struct {
	rooms map[string]*domain.Room
}

func NewInMemoryRoomRepository() *InMemoryRoomRepository {
	return &InMemoryRoomRepository{
		rooms: make(map[string]*domain.Room),
	}
}

func (r *InMemoryRoomRepository) Get(name string) (*domain.Room, error) {
	room := r.rooms[name]
	return room, nil
}
func (r *InMemoryRoomRepository) GetAll() ([]*domain.Room, error) {
	rooms := make([]*domain.Room, len(r.rooms))
	i := 0
	for _, room := range r.rooms {
		rooms[i] = room
		i++
	}
	return rooms, nil
}

func (r *InMemoryRoomRepository) Create(room *domain.Room) error {
	_, ok := r.rooms[room.Name]
	if ok {
		return errors.New("room already exists")
	}

	r.rooms[room.Name] = room
	return nil
}
