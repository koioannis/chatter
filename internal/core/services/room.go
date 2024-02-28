package services

import (
	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/core/domain"
	"github.com/koioannis/chatter/internal/ports/store"
	"github.com/sirupsen/logrus"
)

type RoomService struct {
	repo   store.RoomRepository
	logger *logrus.Logger
}

func NewRoomService(repo store.RoomRepository, logger *logrus.Logger) *RoomService {
	return &RoomService{
		repo:   repo,
		logger: logger,
	}
}

func (s *RoomService) GetAll() []*domain.Room {
	rooms, _ := s.repo.GetAll()
	return rooms
}

func (s *RoomService) Create(name string) (*domain.Room, error) {
	room, err := s.repo.Get(name)
	if err != nil {
		return nil, err
	}

	if room != nil {
		return nil, domain.ErrRoomAlreadyExists
	}

	room, err = domain.NewRoom(uuid.New(), name)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(room); err != nil {
		return nil, err
	}

	return room, nil
}
