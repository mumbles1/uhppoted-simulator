package UT0311L04

import (
	"fmt"

	"codeberg.org/uhppoted/uhppoted-core/messages"
	"codeberg.org/uhppoted/uhppoted-core/types"

	"codeberg.org/uhppoted/uhppoted-simulator/entities"
)

func (s *UT0311L04) setDoorControlState(request *messages.SetDoorControlStateRequest) (*messages.SetDoorControlStateResponse, error) {
	if request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	if request.Door < 1 || request.Door > 4 {
		fmt.Printf("ERROR: Invalid door' - expected: [1..4], received:%d", request.Door)
		return nil, nil
	}

	door := request.Door

	switch request.ControlState {
	case 1:
		s.Doors.SetMode(door, types.ModeNormallyOpen)

	case 2:
		s.Doors.SetMode(door, types.ModeNormallyClosed)

	case 3:
		s.Doors.SetMode(door, types.ModeControlled)
	}

	s.Doors.SetDelay(door, entities.Delay(uint64(request.Delay)*1000000000))

	mode := uint8(0)
	switch s.Doors.Mode(request.Door) {
	case types.ModeNormallyOpen:
		mode = 1
	case types.ModeNormallyClosed:
		mode = 2
	case types.ModeControlled:
		mode = 3
	}

	response := messages.SetDoorControlStateResponse{
		SerialNumber: s.SerialNumber,
		Door:         door,
		ControlState: mode,
		Delay:        s.Doors.Delay(door).Seconds(),
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
