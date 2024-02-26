package store

import (
	"testing"

	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryRoomRepository(t *testing.T) {
	r := NewInMemoryRoomRepository()
	room := domain.NewRoom(uuid.New(), "foo")
	r.Create(room)

	actualRoom, err := r.Get(room.Name)
	assert.Nil(t, err)
	assert.Equal(t, room, actualRoom)

	rooms, err := r.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(rooms))
}
