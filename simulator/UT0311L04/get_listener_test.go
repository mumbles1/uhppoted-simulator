package UT0311L04

import (
	"net/netip"
	"testing"

	"codeberg.org/uhppoted/uhppoted-core/messages"
)

func TestHandleGetListener(t *testing.T) {
	request := messages.GetListenerRequest{
		SerialNumber: 405419896,
	}

	response := messages.GetListenerResponse{
		SerialNumber: 405419896,
		AddrPort:     netip.MustParseAddrPort("10.0.0.10:43210"),
	}

	testHandle(&request, &response, t)
}
