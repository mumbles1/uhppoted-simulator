package UT0311L04

import (
	"testing"

	"codeberg.org/uhppoted/uhppoted-core/messages"
)

func TestHandleDeleteCardsRequest(t *testing.T) {
	request := messages.DeleteCardsRequest{
		SerialNumber: 405419896,
		MagicWord:    0x55aaaa55,
	}

	response := messages.DeleteCardsResponse{
		SerialNumber: 405419896,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}
