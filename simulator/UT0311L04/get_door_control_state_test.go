package UT0311L04

import (
	"testing"

	"codeberg.org/uhppoted/uhppoted-core/messages"
)

func TestHandleGetDoorControlState(t *testing.T) {
	request := messages.GetDoorControlStateRequest{
		SerialNumber: 405419896,
		Door:         2,
	}

	response := messages.GetDoorControlStateResponse{
		SerialNumber: 405419896,
		Door:         2,
		ControlState: 2,
		Delay:        22,
	}

	testHandle(&request, &response, t)
}
