package UT0311L04

import (
	"testing"

	"codeberg.org/uhppoted/uhppoted-core/messages"
)

func TestHandleDeleteCardRequest(t *testing.T) {
	request := messages.DeleteCardRequest{
		SerialNumber: 405419896,
		CardNumber:   192837465,
	}

	response := messages.DeleteCardResponse{
		SerialNumber: 405419896,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}
