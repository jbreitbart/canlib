package can

import (
	"testing"
)

// TestChannelMultiplexWorks will test ChannelMultiplex to ensure input messaegs reach output
func TestChannelMultiplexWorks(t *testing.T) {
	message := &Frame{OID: 555}
	canIn := make(chan *Frame)
	out1 := make(chan *Frame)
	out2 := make(chan *Frame)

	go ChannelMultiplex(canIn, out1, out2)
	canIn <- message
	res1 := <-out1
	res2 := <-out2
	close(canIn)

	if !CompareFrames(res1, message) {
		t.Error("The CAN frames did not match")
	}
	if !CompareFrames(res2, message) {
		t.Error("The CAN frames did not match")
	}
}
