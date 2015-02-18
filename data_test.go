package icmpv4

//go:generate go-bindata -pkg=icmpv4 -o=bindata_test.go test_data

//packets contains test ICMP packets captured from the real world. All of their checksums are correct.
var testPackets [][]byte

func init() {
	testPackets = make([][]byte, 0)
	for _, name := range AssetNames() {
		data, _ := Asset(name)
		testPackets = append(testPackets, data)
	}
}
