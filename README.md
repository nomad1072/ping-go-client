# ping-go-client

# Go installation
1. Install the go binary from the official site. ```https://golang.org/doc/install```
2. Add a few environment variables to your ```~/.bash_profile``` file. ```$GOROOT``` should point to your go installation and ```$GOPATH``` should point to your go workspace.
3. It's recommended that your go workspace be split into 3 directories - ```bin, pkg and src``` and the contents in each of them is self-explanatory. A few articles explain how to set this up(https://www.ardanlabs.com/blog/2016/05/installing-go-and-your-workspace.html).
4. Download the code for this repository by running ```go get -u github.com/nomad1072/ping-go-client```.

# Steps to run the program

1. Clone the repository. RUN ```go get -u github.com/nomad1072/ping-go-client```.
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

<img src="https://mybucket-test-openwhisk.s3.amazonaws.com/ipv4.png"
     alt="Ping IPv4 output" />
     
# IPv6

<img src="https://mybucket-test-openwhisk.s3.amazonaws.com/ipv6.png"
     alt="Ping IPv6 output"/>
