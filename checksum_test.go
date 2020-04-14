package icmpv4

import "testing"

func TestChecksum(t *testing.T) {
	for i, p := range testPackets {
		if c := checksum(p); c != 0 {
			t.Errorf("Packet %d: Expected checksum 0, got %#x", i, c)
		}
		pNew := make([]byte, len(p))
		copy(pNew, p)
		old := uint16(pNew[2])<<8 | uint16(pNew[3])

		pNew[2] = 0
		pNew[3] = 0
		if c := checksum(pNew); c != old {
			t.Errorf("Packet %d: Expected checksum %#x, got %#x", i, old, c)
		}
	}
}
