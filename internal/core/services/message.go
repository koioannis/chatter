package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/core/domain"
	"github.com/koioannis/chatter/internal/dto"
	"github.com/koioannis/chatter/internal/ports/events"
	"github.com/koioannis/chatter/internal/ports/store"
	"github.com/sirupsen/logrus"
)

type MessageService struct {
	messageRepo  store.MessageRepository
	roomRepo     store.RoomRepository
	msgPublisher events.MessagePublisher
	logger       *logrus.Logger
}

func NewMessageService(messageRepo store.MessageRepository, roomRepo store.RoomRepository, msgPublisher events.MessagePublisher, logger *logrus.Logger) *MessageService {
	return &MessageService{
		messageRepo:  messageRepo,
		roomRepo:     roomRepo,
		msgPublisher: msgPublisher,
		logger:       logger,
	}
}

func (m *MessageService) Create(createMessageDto dto.CreateMessageDTO) (*domain.Message, error) {
	room, err := m.roomRepo.GetById(createMessageDto.RoomId())
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
	err = m.messageRepo.Create(message)
	go m.publishMessage(message)
	return message, err
}

func (m *MessageService) publishMessage(msg *domain.Message) {
	fmt.Println("Publishing from sevice")
	if err := m.msgPublisher.Publish(msg); err != nil {
		m.logger.Error(err)
	}
}

func (m *MessageService) GetAllByRoomId(roomId uuid.UUID) ([]*domain.Message, error) {
	return m.messageRepo.GetAllByRoomId(roomId)
}
