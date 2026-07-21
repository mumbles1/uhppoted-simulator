package UT0311L04

import (
	"testing"

	"codeberg.org/uhppoted/uhppoted-core/messages"
)

func TestHandleGetCardsRequest(t *testing.T) {
	request := messages.GetCardsRequest{
		SerialNumber: 405419896,
	}

	response := messages.GetCardsResponse{
		SerialNumber: 405419896,
		Records:      3,
	}

	testHandle(&request, &response, t)
}
