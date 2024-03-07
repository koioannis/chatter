package domain

import "github.com/google/uuid"

type Message struct {
	Id      uuid.UUID
	Content string
	Sender  string
	RoomId  uuid.UUID
}

func NewMessage(id uuid.UUID, senderUsername string, roomID uuid.UUID, content string) (*Message, error) {
	if content == "" {
		return nil, ErrEmptyMessage
	}

	return &Message{
		Id:      id,
		Content: content,
		Sender:  senderUsername,
		RoomId:  roomID,
	}, nil
}
