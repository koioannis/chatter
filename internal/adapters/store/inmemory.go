package store

import (
	"errors"

	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/core/domain"
)

type InMemoryRoomRepository struct {
	rooms map[string]*domain.Room
	index map[uuid.UUID]*domain.Room
}

func NewInMemoryRoomRepository() *InMemoryRoomRepository {
	return &InMemoryRoomRepository{
		rooms: make(map[string]*domain.Room),
		index: make(map[uuid.UUID]*domain.Room),
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
	r.index[room.Id] = room
	return nil
}

func (r *InMemoryRoomRepository) GetById(id uuid.UUID) (*domain.Room, error) {
	room := r.index[id]
	return room, nil
}

func (r *InMemoryRoomRepository) Update(room *domain.Room) error {
	_, ok := r.rooms[room.Name]
	if ok {
		return errors.New("room already exists")
	}

	r.rooms[room.Name] = room
	return nil
}

type InMemoryMessageRepository struct {
	messages map[uuid.UUID][]*domain.Message
}

func NewInMemoryMessageRepository() *InMemoryMessageRepository {
	return &InMemoryMessageRepository{
		messages: make(map[uuid.UUID][]*domain.Message),
	}
}

func (r *InMemoryMessageRepository) Create(message *domain.Message) error {
	if _, ok := r.messages[message.RoomId]; !ok {
		r.messages[message.RoomId] = make([]*domain.Message, 0)
	}

	r.messages[message.RoomId] = append(r.messages[message.RoomId], message)
	return nil
}

func (r *InMemoryMessageRepository) GetAllByRoomId(roomId uuid.UUID) ([]*domain.Message, error) {
	messages := r.messages[roomId]
	return messages, nil
}
