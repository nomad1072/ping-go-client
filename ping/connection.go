package ping

import (
	"strings"

	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"

	"golang.org/x/net/icmp"
)

func ipType(address string) string {
	if strings.Count(address, ":") < 2 {
		return "ipv4"
	} else if strings.Count(address, ":") >= 2 {
		return "ipv6"
	}
	return "invalid"
}

func prepareConnection(address string, ttl int, sourceIpv6 string) (*icmp.PacketConn, string) {
	iptype := ipType(address)
	var connection *icmp.PacketConn

	if iptype == "ipv4" {
		connection, _ = icmp.ListenPacket("udp4", "0.0.0.0")
		connection.IPv4PacketConn().SetControlMessage(ipv4.FlagTTL, true)
		connection.IPv4PacketConn().SetTTL(ttl)
		return connection, iptype
	} else if iptype == "ipv6" {
		connection, _ = icmp.ListenPacket("udp6", sourceIpv6)
		connection.IPv6PacketConn().SetControlMessage(ipv6.FlagHopLimit, true)
		connection.IPv6PacketConn().SetHopLimit(ttl)
		return connection, iptype
	}
	return nil, "invalid"
}
