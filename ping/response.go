package ping

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

func processResponseIPv4(connection *icmp.PacketConn, sequence int) {
	rb := make([]byte, 1500)
	start := time.Now()
	n, peer, _, err := connection.IPv4PacketConn().ReadFrom(rb)
	if err != nil {
		fmt.Println("Request timeout for icmp_seq ", sequence)
		return
	}
	rm, err := icmp.ParseMessage(ipv4.ICMPTypeEcho.Protocol(), rb[:n])
	if err != nil {
		log.Fatal(err)
	}
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		duration := time.Since(start)
		fmt.Printf("Received %v bytes from %v  icmp_seq=%v ttl=%v time=%v", n, peer.Src, sequence, peer.TTL, duration)
		fmt.Println()
	default:
		log.Printf("got %+v; want echo reply", rm.Type)
	}
}

func processResponseIPv6(connection *icmp.PacketConn, sequence int) {
	rb := make([]byte, 1500)
	start := time.Now()
	n, peer, _, err := connection.IPv6PacketConn().ReadFrom(rb)
	if err != nil {
		log.Fatal(err)
	}

	rm, err := icmp.ParseMessage(ipv6.ICMPTypeEchoRequest.Protocol(), rb[:n])
	if err != nil {
		log.Fatal(err)
	}

	switch rm.Type {
	case ipv6.ICMPTypeEchoReply:
		duration := time.Since(start)
		fmt.Printf("Received %v bytes from %v  icmp_seq=%v ttl=%v time=%v", n, peer.Src, sequence, peer.HopLimit, duration)
		fmt.Println()
	default:
		log.Printf("got %+v; want echo reply", rm.Type)
	}
}
