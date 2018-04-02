package canlib

import (
	"testing"
)

// TestCanFrameToString will verify that a CAN message is formatted properly
func TestCanFrameToString(t *testing.T) {
	testFrame := CanFrame{
		OID:       1,
		Data:      []byte{1},
		Timestamp: 1000000000,
	}
	expected := "1,1,NOEFF,NORTR,NOERR,1,1,01"
	result := CanFrameToString(testFrame, ",")
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

// TestProcessedCanFrameToString makes sure that ProcessedCanFrameToString works
func TestProcessedCanFrameToString(t *testing.T) {
	testRawFrame := CanFrame{
		OID:       1,
		Data:      []byte{1},
		Timestamp: 1000000000,
	}
	testProcessedFrame := ProcessedCanFrame{Packet: testRawFrame,
		PacketHash: "testHash"}
	expected := "testHash,1,1,NOEFF,NORTR,NOERR,1,1,01"
	result := ProcessedCanFrameToString(testProcessedFrame, ",")
	if expected != result {
		t.Errorf("%s != %s", expected, result)
	}
}
