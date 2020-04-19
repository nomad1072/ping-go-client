package ping

import (
	"log"
	"net"
	"time"
)

// Ping send an ICMP Echo request to a given host with
// set count specifying number of requests and ttl
// specifies the maximum network hops
func Ping(address string, count int, ttl int, sourceIpv6 string) {
	// Listen for icmp packets
	connection, iptype := prepareConnection(address, ttl, sourceIpv6)
	defer connection.Close()
	println("Ip Type: ", iptype)
	// Resolve IP to hostname
	ipaddress, _ := net.ResolveIPAddr("ip", address)

	// create the destination Ip and DNS Zone to send the package to
	destination := &net.UDPAddr{IP: ipaddress.IP, Zone: ipaddress.Zone}

	// Create an icmp message
	for i := 1; i <= count; i++ {
		wb := prepareMessage(iptype, i+1)
		if _, err := connection.WriteTo(wb, destination); err != nil {
			log.Fatalf("WriteTo err, %s", err)
		}

		if iptype == "ipv4" {
			processResponseIPv4(connection, (i))
		} else if iptype == "ipv6" {
			processResponseIPv6(connection, (i))
		}
		time.Sleep(3000)
	}
}
