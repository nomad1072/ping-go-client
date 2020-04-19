# ping-go-client

# Steps to run the program

1. Clone the repository.
2. ```cd``` into the project directory.
3. RUN ```go test ./... -v```. You should see ok messages from each test case.
4. RUN ```go run main.go``` with some optional command line arguments.
5. Command Line Args: ```-hostname=<ip-or-hostname> -count=<number-of-requests> -ttl=<network-hops>```.
6. Default values for ```count=20 and ttl=255```.
7. If testing for an Ipv6 address also specify ```sourceip=<current-host-ipv6-address>```.

# Features Implemented

1. Setup ttl to restrict the transfer of packets to a set number of network hops.
2. Each request reports the number of bytes received, the destination ip and Round trip time(RTT).
3. Count feature as mentioned in ping man page to restrict the number of icmp requests to this variable.

# Output

# IPv4

https://imgur.com/MpAf1BH

# IPv6

https://imgur.com/gallery/YNs3aSK
