package events

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/core/domain"
)

type InMemoryPublisherSubscriber struct {
	sync.RWMutex
	subscribers map[uuid.UUID][]chan *domain.Message
}

func NewInmemoryPublisherSubscriber() *InMemoryPublisherSubscriber {
	return &InMemoryPublisherSubscriber{
		subscribers: make(map[uuid.UUID][]chan *domain.Message),
	}
}

func (m *InMemoryPublisherSubscriber) Publish(msg *domain.Message) error {
	m.RLock()
	defer m.RUnlock()

	for _, ch := range m.subscribers[msg.RoomId] {
		// we could do a select here, but since we don't have buffered channels
		// it's fine
		ch <- msg
		fmt.Println("Publishing done")
	}

	return nil
}

func (k *InMemoryPublisherSubscriber) Subscribe(roomId uuid.UUID) <-chan *domain.Message {
	k.Lock()
	defer k.Unlock()

	ch := make(chan *domain.Message, 100)
	if _, ok := k.subscribers[roomId]; !ok {
		fmt.Println("making a sub for roomId", roomId)
		k.subscribers[roomId] = make([]chan *domain.Message, 0)
	}
	k.subscribers[roomId] = append(k.subscribers[roomId], ch)
	return ch
}
