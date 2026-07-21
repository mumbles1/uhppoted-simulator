package UT0311L04

import (
	"fmt"
	"time"

	"codeberg.org/uhppoted/uhppoted-core/messages"
	"codeberg.org/uhppoted/uhppoted-core/types"
	"codeberg.org/uhppoted/uhppoted-simulator/entities"
)

func (s *UT0311L04) setTime(request *messages.SetTimeRequest) (*messages.SetTimeResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	dt := time.Time(request.DateTime).Format("2006-01-02 15:04:05")
	if utc, err := time.ParseInLocation("2006-01-02 15:04:05", dt, time.UTC); err != nil {
		return nil, err
	} else {
		now := time.Now().UTC()
		delta := utc.Sub(now)
		datetime := now.Add(delta)

		s.TimeOffset = entities.Offset(delta)
		response := messages.SetTimeResponse{
			SerialNumber: s.SerialNumber,
			DateTime:     types.DateTime(datetime),
		}

		if err = s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		return &response, nil
	}
}
