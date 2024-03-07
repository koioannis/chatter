package dto

import "github.com/google/uuid"

type CreateMessageDTO struct {
	content string
	sender  string
	roomId  uuid.UUID
}

func (dto CreateMessageDTO) Content() string {
	return dto.content
}

func (dto CreateMessageDTO) Sender() string {
	return dto.sender
}

func (dto CreateMessageDTO) RoomId() uuid.UUID {
	return dto.roomId
}
