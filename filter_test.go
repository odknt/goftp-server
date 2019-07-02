package server

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleFilterCheckIP(t *testing.T) {
	// list is empty
	filter := &SimpleFilter{}
	assert.NoError(t, filter.CheckIP(nil, ""))

	filter = &SimpleFilter{
		AllowIPs: []net.IP{
			net.ParseIP("192.168.1.1"),
			net.ParseIP("192.168.1.2"),
		},
	}
	assert.NoError(t, filter.CheckIP(net.ParseIP("192.168.1.1"), ""))
	assert.NoError(t, filter.CheckIP(net.ParseIP("192.168.1.2"), ""))
	assert.Error(t, filter.CheckIP(net.ParseIP("192.168.1.3"), ""))
}
