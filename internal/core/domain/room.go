package domain

import "github.com/google/uuid"

type Room struct {
	Id   uuid.UUID
	Name string
}

func NewRoom(id uuid.UUID, name string) *Room {
	return &Room{
		Id:   id,
		Name: name,
	}
}
