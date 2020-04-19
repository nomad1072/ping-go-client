package ping

import (
	"log"
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

func prepareMessage(iptype string, sequence int) []byte {
	var icmpRequestType icmp.Type
	if iptype == "ipv4" {
		icmpRequestType = ipv4.ICMPTypeEcho
	} else if iptype == "ipv6" {
		icmpRequestType = ipv6.ICMPTypeEchoRequest
	}

	message := &icmp.Message{
		Type: icmpRequestType, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: sequence,
			Data: []byte("ECHO-REQUEST-ECHO-REQUEST"),
		},
	}

	wb, err := message.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}

	return wb
}
