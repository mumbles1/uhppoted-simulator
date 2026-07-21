package UT0311L04

import (
	"codeberg.org/uhppoted/uhppoted-core/messages"
	"codeberg.org/uhppoted/uhppoted-core/types"
)

func (s *UT0311L04) getDevice(request *messages.GetDeviceRequest) (*messages.GetDeviceResponse, error) {
	if request.SerialNumber != 0 && request.SerialNumber != s.SerialNumber {
		return nil, nil
	}

	response := messages.GetDeviceResponse{
		SerialNumber: s.SerialNumber,
		IpAddress:    s.IpAddress,
		SubnetMask:   s.SubnetMask,
		Gateway:      s.Gateway,
		MacAddress:   s.MacAddress,
		Version:      types.Version(s.Version),
		Date:         s.Released,
	}

	return &response, nil
}
