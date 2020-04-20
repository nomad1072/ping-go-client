package ping

import (
	"testing"
)

func TestConnection(t *testing.T) {
	connection, iptype := prepareConnection("amazon.com", 255)
	connection1, iptype1 := prepareConnection("2001:4860:4860::8888", 255)
	t.Run("connection=typeicmp", func(t *testing.T) {
		if iptype != "ipv4" {
			t.Errorf("Test failed, invalid connection object. expected ipv4 but got %s", iptype)
		}
	})
	t.Run("connection=typeicmpipv6", func(t *testing.T) {
		if iptype1 != "ipv6" {
			t.Errorf("Test failed, invalid connection object. expected ipv4 but got %s", iptype1)
		}
	})

	t.Run("connection=typeicmpNotNil", func(t *testing.T) {
		if connection == nil {
			t.Errorf("Test failed, invalid connection object. expected of type icmp.PacketConn but got nil")
		}
	})
	t.Run("connection=typeicmpipv6NotNil", func(t *testing.T) {
		if connection1 == nil {
			t.Errorf("Test failed, invalid connection object. expected icmp.PacketConn but got nil")
		}
	})

}
