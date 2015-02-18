package icmpv4

import "testing"

func TestHeaderOptionsByte(t *testing.T) {
	for pi, p := range testPackets {
		packet, err := Parse(p)
		if err != nil {
			t.Errorf("Packet %d: Error parsing packet: %s", pi, err)
			continue
		}

		h := packet.HeaderOptions
		var i uint8
		for i = 0; i < 4; i++ {
			if h.Byte(i) != p[i+4] {
				t.Errorf("Packet %d: Expected: %#x, Got Byte(%d) = %#x", pi, p[i+4], i, h.Byte(i))
			}
			h.SetByte(i, i)
			if h.Byte(i) != i {
				t.Errorf("Packet %d: SetByte(%d, %d): Expected: %#x, Got Byte(%d) = %#x", pi, i, i, i, i, h.Byte(i))
			}
		}

	}
}

func TestHeaderOptionsUnint16(t *testing.T) {
	for pi, p := range testPackets {
		packet, err := Parse(p)
		if err != nil {
			t.Errorf("Packet %d: Error parsing packet: %s", pi, err)
			continue
		}

		h := packet.HeaderOptions
		var i uint8
		for i = 0; i < 2; i++ {
			if h.Uint16(i) != uint16(p[2*i+4])<<8|uint16(p[2*i+5]) {
				t.Errorf("Packet %d: Expected: %#x, Got Uint16(%d) = %#x", pi, uint16(p[i+4])<<8|uint16(p[i+5]), i, h.Uint16(i))
			}
			h.SetUint16(i, uint16(i))
			if h.Uint16(i) != uint16(i) {
				t.Errorf("Packet %d: SetUint16(%d, %d): Expected: %#x, Got Uint16(%d) = %#x", pi, i, i, i, i, h.Uint16(i))
			}
		}

	}
}

func TestMalformedHeaders(t *testing.T) {
	// malformed
	_, err := Parse([]byte{})
	if _, ok := err.(InvalidPacketError); !ok {
		t.Errorf("Expected InvalidPacketError, Got: %#v", err)

	}
	if err.Error() != "Malformed headers" {
		t.Errorf("Expected \"Malformed headers\" error, Got: %s", err.Error())
	}

	_, err = Parse([]byte{1, 2, 3, 4})
	if _, ok := err.(InvalidPacketError); !ok {
		t.Errorf("Expected InvalidPacketError, Got: %#v", err)
	}
	if err.Error() != "Malformed headers" {
		t.Errorf("Expected \"Malformed headers\" error, Got: %s", err.Error())
	}

	// bad checksum
	_, err = Parse([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if _, ok := err.(InvalidPacketError); !ok {
		t.Errorf("Expected InvalidPacketError, Got: %#v", err)
	}
	if err.Error() != "Invalid checksum" {
		t.Errorf("Expected \"Invalid checksum\" error, Got: %s", err.Error())
	}
}

func TestParseMarshal(t *testing.T) {
	for pi, p := range testPackets {
		packet, err := Parse(p)
		if err != nil {
			t.Errorf("Packet %d: Error parsing packet: %s", pi, err)
			continue
		}
		//check checksums work correctly
		b := packet.Marshal()
		if checksum(b) != 0 {
			t.Errorf("Packet %d: Expected checksum %#x, got %#x", pi, 0, checksum(b))
			continue
		}
		b[2] = 0
		b[3] = 0
		if checksum(b) != packet.Checksum {
			t.Errorf("Packet %d: Expected checksum %#x, got %#x", pi, packet.Checksum, checksum(b))
			continue
		}

		//reset
		b = packet.Marshal()

		//check data matches
		if len(p) != len(b) {
			t.Errorf("Packet %d: Expected len(p): %d, got %d", pi, len(p), len(b))
			continue
		}
		for i := range p {
			if p[i] != b[i] {
				t.Errorf("Packet %d: packet data mismatch at p[%d]", pi, i)
				break
			}
		}
	}
}
