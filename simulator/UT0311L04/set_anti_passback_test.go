package UT0311L04

import (
	"testing"

	"codeberg.org/uhppoted/uhppoted-core/messages"
)

func TestHandleSetAntiPassback(t *testing.T) {
	request := messages.SetAntiPassbackRequest{
		SerialNumber: 405419896,
		AntiPassback: 0x04,
	}

	response := messages.SetAntiPassbackResponse{
		SerialNumber: 405419896,
		Ok:           true,
	}

	testHandle(&request, &response, t)
}
