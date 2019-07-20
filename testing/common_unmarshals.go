package ldetesting

import (
	"net"
	"strconv"
	"time"
)

func (p *Custom) unmarshalTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

func (p *Custom) unmarshalAddrIP(s string) (net.IP, error) {
	return net.ParseIP(s), nil
}

func (p *CustomBuiltin) unmarshalField(s string) (int, error) {
	return strconv.Atoi(s)
}
