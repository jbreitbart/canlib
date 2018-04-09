package can

import (
	"testing"
)

// TestFrameToString will verify that a CAN message is formatted properly
func TestFrameToString(t *testing.T) {
	testFrame := Frame{
		OID:       1,
		Data:      []byte{1},
		Timestamp: 1000000000,
	}
	expected := "1,1,NOEFF,NORTR,NOERR,1,1,01"
	result := testFrame.ToString(",")
	if expected != result {
		t.Errorf("%s != %s", expected, result)
	}
}

// TestTimestampToSeconds makes sure that the function works
func TestTimestampToSeconds(t *testing.T) {
	fakeTime := int64(1000000000)
	expected := float64(1)
	result := TimestampToSeconds(fakeTime)
	if expected != result {
		t.Errorf("%v != %v", expected, result)
	}
}

// TestProcessedFrameToString makes sure that ProcessedFrameToString works
func TestProcessedFrameToString(t *testing.T) {
	testRawFrame := Frame{
		OID:       1,
		Data:      []byte{1},
		Timestamp: 1000000000,
	}
	testProcessedFrame := ProcessedFrame{Packet: testRawFrame,
		PacketHash: "testHash"}
	expected := "testHash,1,1,NOEFF,NORTR,NOERR,1,1,01"
	result := testProcessedFrame.ToString(",")
	if expected != result {
		t.Errorf("%s != %s", expected, result)
	}
}
