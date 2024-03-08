package services

import (
	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/core/domain"
	"github.com/koioannis/chatter/internal/dto"
	"github.com/koioannis/chatter/internal/ports/store"
)

type MessageService struct {
	message_repo store.MessageRepository
	room_repo    store.RoomRepository
}

func NewMessageService(message_repo store.MessageRepository, room_repo store.RoomRepository) *MessageService {
	return &MessageService{
		message_repo: message_repo,
		room_repo:    room_repo,
	}
}

func (m *MessageService) Create(createMessageDto dto.CreateMessageDTO) (*domain.Message, error) {
	room, err := m.room_repo.GetById(createMessageDto.RoomId())
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, domain.ErrRoomDoesNotExist

	}

	message, err := domain.NewMessage(uuid.New(), createMessageDto.Sender(), createMessageDto.RoomId(), createMessageDto.Content())
	if err != nil {
		return nil, err
	}

	err = m.message_repo.Create(message)
	return message, err
}

func (m *MessageService) GetAllByRoomId(roomId uuid.UUID) ([]*domain.Message, error) {
	return m.message_repo.GetAllByRoomId(roomId)
}
