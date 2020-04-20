package main

import (
	"flag"
	"fmt"
	"math"
	"net"

	"github.com/nomad1072/ping-go-client/ping"
)

func validateArgs(count int, ip string) (string, bool) {
	ipaddress, _ := net.ResolveIPAddr("ip", ip)
	if ipaddress == nil {
		return "invalid IP address/hostname provided", false
	}

	if count < 0 {
		return "count accepts positive integers", false
	}

	return "All arguments valid", true
}

func fetchArgs() (string, int, int, string) {
	var ip = flag.String("hostname", "cloudflare.com", "The ip or the hostname")
	var count = flag.Int("count", math.MaxInt64, "Number of requests to be made")
	var ttl = flag.Int("ttl", 255, "Number of maximum network hops")
	var sourceIP = flag.String("sourceip", "2601:281:8380:3570:f4db:2ca7:3f8f:a109", "Specify source IP for IPv6")

	flag.Parse()
	// var address = "2001:4860:4860::8888"
	fmt.Println("Ip: ", *ip)
	fmt.Println("Count: ", *count)
	return *ip, *count, *ttl, *sourceIP
}

func main() {
	ip, count, ttl, sourceIpv6 := fetchArgs()
	message, ok := validateArgs(count, ip)
	println("Message: ", message)
	if ok {
		ping.Ping(ip, count, ttl, sourceIpv6)
	}
}
