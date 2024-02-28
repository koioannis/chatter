package domain

import "github.com/google/uuid"

type Room struct {
	Id   uuid.UUID
	Name string
}

func NewRoom(id uuid.UUID, name string) (*Room, error) {
	if name == "" {
		return nil, ErrInvalidRoomName
	}

	if len(name) <= 2 {
		return nil, ErrInvalidRoomName
	}

	return &Room{
		Id:   id,
		Name: name,
	}, nil
}
