package simulator

import (
	"time"

	"codeberg.org/uhppoted/uhppoted-core/messages"
	"codeberg.org/uhppoted/uhppoted-core/types"

	"codeberg.org/uhppoted/uhppoted-simulator/entities"
)

type Simulator interface {
	DeviceID() uint32
	DeviceType() string
	FilePath() string

	Handle(messages.Request) (any, error)
	Tick()
	RunTasks()
	Save() error
	Delete() error

	Swipe(cardNumber uint32, door uint8, direction entities.Direction, PIN uint32) (bool, error)
	Passcode(door uint8, passcode uint32) (bool, error)
	Open(door uint8) (bool, error)
	Close(door uint8) (bool, error)
	ButtonPressed(door uint8, duration time.Duration) (bool, error)
	StoreCard(card uint32, from types.Date, to types.Date, doors []uint8, PIN uint32) error
	Reset() (bool, error)
}
