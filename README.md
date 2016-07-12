# icmpv4
--
    import "github.com/korylprince/go-icmpv4"

Package icmpv4 provides functions similar to the standard library's net package
for working with ICMPv4 generic packets. See sub-package
github.com/korylprince/go-icmpv4/echo for an example of using this package to
ping IPs.

Note that in most cases you will need to run the compiled binary with
root/administrator privileges in order to send or receive ICMPv4 packets.

## Usage

```go
const ICMPv4HeaderLength = 8
```
Length of ICMPv4 header in bytes

#### func  Dial

```go
func Dial(laddr, raddr *net.IPAddr) (*net.IPConn, error)
```
Dial creates an ICMPv4 *net.IPconn from laddr to raddr

#### func  Listen

```go
func Listen(laddr *net.IPAddr) (*net.IPConn, error)
```
Listen listens for incoming ICMPv4 packets

#### func  Listener

```go
func Listener(conn *net.IPConn, packets chan *IPPacket, errors chan error, done chan struct{}) error
```
Listener parses incoming ICMPv4 packets from an ICMPv4 net.IPConn and sends
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
func Send(laddr, raddr *net.IPAddr, packet []byte) (err error)
```
Send sends an ICMPv4 packet from laddr to raddr

#### type HeaderOptions

```go
type HeaderOptions uint32
```

HeaderOptions is a 4-byte part of the ICMPv4 header. The value(s) it represents
depends on the Type and Code.

#### func (HeaderOptions) Byte

```go
func (h HeaderOptions) Byte(n uint8) byte
```
Byte returns the nth byte of the HeaderOptions, indexed by 0. [0,3] is the only
acceptable range.

#### func (*HeaderOptions) SetByte

```go
func (h *HeaderOptions) SetByte(n uint8, b byte)
```
SetByte sets the nth byte of the HeaderOptions to b, indexed by 0. [0,3] is the
only acceptable range.

#### func (*HeaderOptions) SetUint16

```go
func (h *HeaderOptions) SetUint16(n uint8, i uint16)
```
SetUint16 sets the nth uint16 of the HeaderOptions to i, indexed by 0. [0,1] is
the only acceptable range.

#### func (HeaderOptions) Uint16

```go
func (h HeaderOptions) Uint16(n uint8) uint16
```
Uint16 returns the nth uint16 of the HeaderOptions, indexed by 0. [0,1] is the
only acceptable range.

#### type IPPacket

```go
type IPPacket struct {
	*Packet
	LocalAddr  *net.IPAddr
	RemoteAddr *net.IPAddr
}
```

IPPacket is a wrapper for Packet with IP information

#### type InvalidPacketError

```go
type InvalidPacketError string
```

InvalidPacketError denotes and error parsing an ICMPv4 packet

#### func (InvalidPacketError) Error

```go
func (err InvalidPacketError) Error() string
```

#### type Packet

```go
type Packet struct {
	Type          uint8
	Code          uint8
	Checksum      uint16
	HeaderOptions HeaderOptions
	Body          []byte
}
```

Packet represents an ICMPv4 packet

#### func  Parse

```go
func Parse(b []byte) (*Packet, error)
```
Parse parses a raw ICMPv4 packet and returns an Packet, or an error if one
occurred

#### func (*Packet) Marshal

```go
func (p *Packet) Marshal() []byte
```
Marshal creates a raw ICMPv4 packet from an Packet
