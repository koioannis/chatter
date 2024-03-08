package dto

import "github.com/google/uuid"

type CreateMessageDTO struct {
	content string
	sender  string
	roomId  uuid.UUID
}

func NewCreateMessageDTO(content string, sender string, roomId uuid.UUID) CreateMessageDTO {
	return CreateMessageDTO{
		content: content,
		sender:  sender,
		roomId:  roomId,
	}
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
