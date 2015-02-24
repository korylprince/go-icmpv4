package icmpv4

import (
	"net"
	"testing"
	"time"
)

func TestBadAddr(t *testing.T) {
	laddr, err := net.ResolveIPAddr("ip4", "8.8.8.8")
	if err != nil {
		t.Errorf("Cannot resolve 8.8.8.8 to net.IPAddr")
		return
	}

	//send request
	p := Packet{HeaderOptions: HeaderOptions(0xabad1dea)}
	err = Send(laddr, laddr, p.Marshal())
	if err == nil {
		t.Errorf("Did not receive error when using invalid local address")
		return
	}
}

func pingSetup(t *testing.T, packet []byte) (packets chan *IPPacket, errors chan error) {
	//set up listener
	laddr, err := net.ResolveIPAddr("ip4", "localhost")
	if err != nil {
		t.Errorf("Cannot resolve localhost")
		return
	}
	conn, err := Listen(laddr)
	if err != nil {
		t.Errorf("Cannot listen on localhost")
		return
	}

	packets = make(chan *IPPacket, 10)
	errors = make(chan error, 10)
	done := make(chan struct{})
	retErr := make(chan error)
	go func() {
		err := Listener(conn, packets, errors, done)
		if err != nil {
			retErr <- err
			return
		}
		retErr <- nil
	}()

	//send request
	err = Send(nil, laddr, packet)
	if err != nil {
		t.Errorf("Cannot send on localhost")
		return
	}
	//wait a bit
	time.Sleep(10 * time.Millisecond)

	//signal listener to stop
	close(done)

	//end any pending conn.ReadFromIP's
	err = conn.SetDeadline(time.Now())
	if err != nil {
		t.Errorf("Error setting timeout")
		return
	}

	//wait a bit more
	time.Sleep(10 * time.Millisecond)

	//check that listener stopped correctly
	select {
	case err := <-retErr:
		if err != nil {
			t.Errorf("Error closing Listener: %s", err)
			return
		}
	default:
		t.Errorf("Listener did not stop")
		return
	}
	return packets, errors
}

func TestGoodPing(t *testing.T) {
	p := &Packet{HeaderOptions: HeaderOptions(0xabad1dea)}
	packets, errors := pingSetup(t, p.Marshal())

	//check that no errors happened
errorloop:
	for {
		select {
		case err := <-errors:
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}
				t.Errorf("Received error: %s", err)
				return
			}
		default:
			break errorloop
		}
	}

	//make sure we got the right packet
	select {
	case packet := <-packets:
		if packet == nil {
			t.Errorf("Did not receive packet")
			return
		}
		s := p.Marshal()
		r := packet.Marshal()
		if len(s) != len(r) {
			t.Errorf("send and received packet differ in size: len(sent) = %d, len(received) = %d", len(s), len(r))
			return
		}
		if packet.HeaderOptions != 0xabad1dea {
			t.Errorf("send and received packet HeaderOptions differ: sent = %#x, received = %#x", p.HeaderOptions, packet.HeaderOptions)
			return
		}
		for i := range s {
			if s[i] != r[i] {
				t.Errorf("sent packet is not equal to received packet")
				return
			}
		}
	default:
		t.Errorf("Did not receive packet")
	}
}

func TestBadPing(t *testing.T) {
	packets, errors := pingSetup(t, []byte{0, 0, 255, 254, 0, 0, 0, 0})

	//check that we got invalid packet error
errorloop:
	for {
		select {
		case err := <-errors:
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}
				if _, ok := err.(InvalidPacketError); !ok {
					t.Errorf("Expected error: InvalidPacketError. Got: %#v", err)
					return
				}
			}
		default:
			break errorloop
		}
	}

	//make sure packet didn't come through
	select {
	case packet := <-packets:
		t.Errorf("Got packet and should not have: %#v", packet)
	default:
	}
}

func TestListenerAll(t *testing.T) {
	packets := make(chan *IPPacket, 10)
	errors := make(chan error, 10)
	done := make(chan struct{})
	list, err := ListenerAll(packets, errors, done)
	if err != nil {
		t.Errorf("Error getting interfaces: %#v", err)
		return
	}

	if len(list) == 0 {
		t.Errorf("No interfaces returned")
		return
	}

	p := &Packet{HeaderOptions: HeaderOptions(0xabad1dea)}

	check := make(map[string]struct{})

	for _, laddr := range list {
		check[laddr.String()] = struct{}{}
		//send request
		err = Send(nil, laddr, p.Marshal())
		if err != nil {
			t.Errorf("Cannot send packet to: %s", laddr)
			return
		}
	}
	//wait a bit
	time.Sleep(10 * time.Millisecond)

	//signal listener to stop
	close(done)

	//wait a bit longer
	time.Sleep(10 * time.Millisecond)

	//check that no errors happened
errorloop:
	for {
		select {
		case err := <-errors:
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}
				t.Errorf("Received error: %s", err)
				return
			}
		default:
			break errorloop
		}
	}

	//make sure we got the right packets
packetloop:
	for {
		select {
		case packet := <-packets:
			if packet == nil {
				t.Errorf("Did not receive packet")
				return
			}
			s := p.Marshal()
			r := packet.Marshal()
			if len(s) != len(r) {
				t.Errorf("send and received packet differ in size: len(sent) = %d, len(received) = %d", len(s), len(r))
				return
			}
			if packet.HeaderOptions != 0xabad1dea {
				t.Errorf("send and received packet HeaderOptions differ: sent = %#x, received = %#x", p.HeaderOptions, packet.HeaderOptions)
				return
			}
			for i := range s {
				if s[i] != r[i] {
					t.Errorf("sent packet is not equal to received packet")
					return
				}
			}
			if _, ok := check[packet.RemoteAddr.String()]; ok {
				delete(check, packet.RemoteAddr.String())
			}
		default:
			//if we didn't get a packet for every interface
			if len(check) != 0 {
				t.Errorf("Did not receive packet")
			}
			break packetloop
		}
	}
}
