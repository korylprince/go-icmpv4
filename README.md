[![GoDoc](https://godoc.org/gopkg.in/korylprince/go-icmpv4.v1?status.svg)](https://godoc.org/gopkg.in/korylprince/go-icmpv4.v1)

# About

`go-icmpv4` is a library for working with ICMPv4 packets. The `echo` helper library is useful for dealing with ICMPv4 Echo Request/Reply (ping) packets.

# Installing

`go get gopkg.in/korylprince/go-icmpv4.v1`

If you have any issues or questions [create an issue](https://github.com/korylprince/go-icmpv4/issues).

# Usage

`godoc gopkg.in/korylprince/go-icmpv4.v1`

`godoc gopkg.in/korylprince/go-icmpv4.v1/echo`

Example:

```go
package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"gopkg.in/korylprince/go-icmpv4.v1/echo"
)

func printer(in <-chan *echo.IPPacket) {
	for {
		fmt.Println("Response from:", (<-in).RemoteAddr.String())
	}
}

func errPrinter(in <-chan error) {
	for {
		fmt.Printf("%#v\n", <-in)
	}
}

func main() {
	//set up channels
	packets := make(chan *echo.IPPacket)
	go printer(packets)
	errors := make(chan error)
	go errPrinter(errors)
	done := make(chan struct{})

	//start listener
	intList, err := echo.ListenerAll(packets, errors, done)
	if err != nil {
		panic(err)
	}
	for _, intfc := range intList {
		fmt.Println("Listening on:", intfc)
	}

	//send pings to all IPs on subnet
	for i := 1; i < 255; i++ {
		raddr, err := net.ResolveIPAddr("ip4", "192.168.100."+strconv.Itoa(i))
		if err != nil {
			panic(err)
		}
		err = echo.Send(nil, raddr, 0x1234, 1)
		if err != nil {
			panic(err)
		}
	}

	//wait to receive replies
	time.Sleep(5 * time.Second)

	//shut down listener
	close(done)
}
```

# Testing

`go test -v`

Note: testing will need to be run with raw socket privileges (i.e. with `sudo`.)
