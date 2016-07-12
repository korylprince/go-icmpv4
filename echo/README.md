# echo
--
    import "github.com/korylprince/go-icmpv4/echo"

Package echo provides a thin wrapper over github.com/korylprince/go-icmpv4
solely for ICMPv4 Echo Request/Reply packets. This package makes it easy to add
ICMPv4 pings to your program. This example will ping every IP on a subnet and
print responses as they come in:

    package main

    import (
    	"fmt"
    	"net"
    	"strconv"
    	"time"

    	"github.com/korylprince/go-icmpv4/echo"
    )

    func Printer(in chan *echo.IPPacket) {
    	for {
    		fmt.Println("Response from:", (<-in).RemoteAddr.String())
    	}
    }

    func Errors(in chan error) {
    	for {
    		fmt.Printf("%#v\n", <-in)
    	}
    }

    func main() {
    	//set up channels
    	packets := make(chan *echo.IPPacket)
    	go Printer(packets)
    	errors := make(chan error)
    	go Errors(errors)
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

## Usage

#### func  Listener

```go
func Listener(conn *net.IPConn, packets chan *IPPacket, errors chan error, done chan struct{}) error
```
Listener parses incoming ICMPv4 Echo Replys from an ICMPv4 net.IPConn and sends
packets and errors back on channels. When done is closed, it returns an error
(or nil) from conn.Close().

#### func  ListenerAll

```go
func ListenerAll(packets chan *IPPacket, errors chan error, done chan struct{}) ([]*net.IPAddr, error)
```
ListenerAll creates a Listener for all IPv4 connections available. It returns a
list of addresses that it's listening on or an error if it can't get that list.

#### func  Send

```go
func Send(laddr, raddr *net.IPAddr, identifier, sequence uint16) (err error)
```
Send sends an ICMPv4 Echo Request to raddr from laddr with the given identifier
and sequence

#### type IPPacket

```go
type IPPacket struct {
	*Packet
	RemoteAddr *net.IPAddr
	LocalAddr  *net.IPAddr
}
```

IPPacket is a wrapper for Packet with IP information

#### type Packet

```go
type Packet struct {
	*icmpv4.Packet
}
```

Packet represents an ICMPv4 Echo Request/Reply

#### func  NewEchoRequest

```go
func NewEchoRequest(identifier, sequence uint16) *Packet
```
NewEchoRequest creates a new ICMPv4 Echo Request with the given identifier and
sequence

#### func (*Packet) Identifier

```go
func (e *Packet) Identifier() uint16
```
Identifier gets the ICMPv4 Echo Request/Reply identifier

#### func (*Packet) Sequence

```go
func (e *Packet) Sequence() uint16
```
Sequence gets the ICMPv4 Echo Request/Reply sequence

#### func (*Packet) SetIdentifier

```go
func (e *Packet) SetIdentifier(i uint16)
```
SetIdentifier sets the ICMPv4 Echo Request/Reply identifier

#### func (*Packet) SetSequence

```go
func (e *Packet) SetSequence(i uint16)
```
SetSequence sets the ICMPv4 Echo Request/Reply sequence
