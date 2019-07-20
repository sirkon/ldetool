package ldetesting

import (
	"net"
	"time"
)

func (p *Custom) unmarshalTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

func (p *Custom) unmarshalAddrIP(s string) (net.IP, error) {
	return net.ParseIP(s), nil
}
