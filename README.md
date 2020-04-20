# ping-go-client

# Features Implemented

- Each request reports the number of bytes received, the destination ip and Round trip time(RTT).
- **(Extra Credit)** Setup ttl to restrict the transfer of packets to a set number of network hops.
- **(Extra Credit)** Support for both IPv4 and IPv6 addresses.
- **(Extra Credit)** Count feature as mentioned in ping man page to restrict the number of icmp requests to be made against the remote host to a fixed number.

# Important installation instructions

1. On linux had to run - ```sudo sysctl -w net.ipv4.ping_group_range="0   2147483647"``` because I'm using UDP as transport.
2. On Mac, disable csrutil since the new OS(El Capitan and beyond) ship with System Integrity Protection(SIP). This can only be done when you restart your mac in recovery mode(hold the command+R key when you restart). Once you are in recovery mode open the terminal app and run ```csrutil disable```. Please remember to enable this back ```csrutil enable```.

# Instructions to clone

1. ```cd $GOPATH/src``` (Note: See here for environment variables set up, and for general [Go installation instructions](#go-installation))
2. ```mkdir -p github.com/nomad1072```
3. ```cd github.com/nomad1072```
4. ```git clone https://github.com/nomad1072/ping-go-client.git```

# Instruction to test

1. ```cd $GOPATH/src/github.com/nomad1072/ping-go-client ```
2. Fetch  the necessary dependencies by running: ```go get -u ./...```
3. Run tests: ```go test ./... -v```

# Command Line Arguments

- Command Line Args: ```go run main.go -hostname=<ip-or-hostname> -count=<number-of-requests> -ttl=<network-hops>```.
- ```hostname``` is a mandatory argument. ```count``` and ```ttl``` are optional arguments.
- Sample request - 1. ```go run main.go -hostname=cloudflare.com```
- Sample request - 2. ```go run main.go -hostname=cloudflare.com -ttl=57 -count=10``` 
- Sample request - 3. ```go run main.go -hostname=2001:4860:4860::8888```

# Output

# IPv4

<img src="https://mybucket-test-openwhisk.s3.amazonaws.com/ipv4.png"
     alt="Ping IPv4 output" />
     
# IPv6

<img src="https://mybucket-test-openwhisk.s3.amazonaws.com/ipv6.png"
     alt="Ping IPv6 output"/>
     
# Go installation
1. Install the go binary from the official site. ```https://golang.org/doc/install```
2. Add the ```GOROOT``` and ```GOPATH``` environment variables to your ```~/.bash_profile``` file. 
- ```GOROOT``` should point to your go installation which is ```/usr/local/go``` 
- ```GOPATH``` which should point to the workspace that the project will use, to store the sources and the dependencies. Inside the directory represented by GOPATH, make sure you have the following three directories: bin, pkg, and src. 
