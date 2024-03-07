package domain

import (
	"errors"
)

var ErrInvalidRoomName = errors.New("invalid room name")
var ErrRoomAlreadyExists = errors.New("room already exists")
var ErrRoomDoesNotExist = errors.New("room does not exist")
var ErrEmptyMessage = errors.New("empty message")
