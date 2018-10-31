package server

import (
	"errors"
	"net"
)

// Filter the interface provides filtering access by ip address.
type Filter interface {
	CheckIP(net.IP, string) error
}

// SimpleFilter simply implements Filter.
type SimpleFilter struct {
	AllowIPs []net.IP
}

// CheckIP returns nil if the ip address included in list. Else returns error.
func (f *SimpleFilter) CheckIP(ip net.IP, user string) error {
	if f.AllowIPs == nil {
		return nil
	}
	for _, allow := range f.AllowIPs {
		if allow.Equal(ip) {
			return nil
		}
	}
	return errors.New("access denied")
}
