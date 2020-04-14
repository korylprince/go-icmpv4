[![pkg.go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/korylprince/go-icmpv4/v2)

# About

`go-icmpv4` is a library for working with ICMPv4 packets. The `echo` helper library is useful for dealing with ICMPv4 Echo Request/Reply (ping) packets.

# Installing

Using Go Modules:

`go get github.com/korylprince/go-icmpv4/v2`

Using gopkg.in:

`go get gopkg.in/korylprince/go-icmpv4.v2`

# Usage

**Example:**

```go
package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/korylprince/go-icmpv4/v2/echo"
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

# Issues

If you have any issues or questions [create an issue](https://github.com/korylprince/go-icmpv4/issues).

# Testing

`go test -v`

Note: testing will need to be run with raw socket privileges (i.e. with `sudo`.)
