package domain

import (
	"errors"
)

var ErrInvalidRoomName = errors.New("invalid room name")
var ErrRoomAlreadyExists = errors.New("room already exists")
