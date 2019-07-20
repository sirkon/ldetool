package ldetesting

import (
	"fmt"
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

func (p *Boolean) unmarshalCheck(s string) (bool, error) {
	switch s {
	case "0":
		return false, nil
	case "1":
		return true, nil
	default:
		return false, fmt.Errorf("value to unmarshal can only be 0 or 1, got `%s`", s)
	}
}
