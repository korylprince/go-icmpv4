package icmpv4

import (
	"io/ioutil"
	"log"
)

//packets contains test ICMP packets captured from the real world. All of their checksums are correct.
var testPackets [][]byte

func init() {
	testPackets = make([][]byte, 0)
	files, err := ioutil.ReadDir("test_data")
	if err != nil {
		log.Fatalln("Unable to read test_data directory:", err)
	}
	for _, fi := range files {
		buf, err := ioutil.ReadFile("test_data/" + fi.Name())
		if err != nil {
			log.Fatalf("Unable to read %s: %v\n", fi.Name(), err)
		}
		testPackets = append(testPackets, buf)
	}
}
