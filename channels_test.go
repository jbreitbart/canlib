package canlib

import (
	"testing"
)

// TestRawCanChannelMultiplexWorks will test RawCanChannelMultiplex to ensure input messaegs reach output
func TestRawCanChannelMultiplexWorks(t *testing.T) {
	message := CanFrame{OID: 555}
	res1 := CanFrame{}
	res2 := CanFrame{}
	canIn := make(chan CanFrame)
	out1 := make(chan CanFrame)
	out2 := make(chan CanFrame)

	go RawCanChannelMultiplex(canIn, out1, out2)
	canIn <- message
	res1 = <-out1
	res2 = <-out2
	close(canIn)

	if !CompareRawFrames(res1, message) {
		t.Error("The CAN frames did not match")
	}
	if !CompareRawFrames(res2, message) {
		t.Error("The CAN frames did not match")
	}
}
