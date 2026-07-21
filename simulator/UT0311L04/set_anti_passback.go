package UT0311L04

import (
	"fmt"

	"codeberg.org/uhppoted/uhppoted-core/messages"
)

func (s *UT0311L04) setAntiPassback(request *messages.SetAntiPassbackRequest) (*messages.SetAntiPassbackResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	ok := s.AntiPassback.Set(request.AntiPassback)

	response := messages.SetAntiPassbackResponse{
		SerialNumber: s.SerialNumber,
		Ok:           ok,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
