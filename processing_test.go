package canlib

import (
	"bytes"
	"testing"
)

// TestByteArrayToCanFrame checks that ByteArrayToCanFrame accurately converts an Extended CAN frame into a RawCanFrame
func TestByteArrayToCanFrame(t *testing.T) {
	frame := []byte{109, 237, 19, 137, 8, 0, 0, 0, 15, 234, 197, 79, 101, 147, 251, 118}
	expected := RawCanFrame{
		OID:  2299784557,
		ID:   152300909,
		Rtr:  false,
		Err:  false,
		Eff:  true,
		Dlc:  8,
		Data: []byte{15, 234, 197, 79, 101, 147, 251, 118},
	}
	var result = new(RawCanFrame)
	ByteArrayToCanFrame(frame, result, 0)
	if result.OID != expected.OID {
		t.Error("OID mismatch")
	} else if result.ID != expected.ID {
		t.Error("ID mismatch")
	} else if result.Rtr != expected.Rtr {
		t.Error("RTR mismatch")
	} else if result.Err != expected.Err {
		t.Error("ERR mismatch")
	} else if result.Eff != expected.Eff {
		t.Error("EFF mismatch")
	} else if result.Dlc != expected.Dlc {
		t.Error("data length mismatch")
	} else if bytes.Equal(result.Data, expected.Data) != true {
		t.Error("data value mismatch")
	}
}

// BenchmarkByteArrayToCanFrame runs benchmarks on ByteArrayToCanFrame
func BenchmarkByteArrayToCanFrame(b *testing.B) {
	frame := []byte{109, 237, 19, 137, 8, 0, 0, 0, 15, 234, 197, 79, 101, 147, 251, 118}
	var result = new(RawCanFrame)
	b.ResetTimer()
	ByteArrayToCanFrame(frame, result, 0)
	for i := 0; i < b.N; i++ {
		ByteArrayToCanFrame(frame, result, 0)
	}
}

// TestProcessRawCan will verify that can messages can be processed
func TestProcessRawCan(t *testing.T) {
	testFrame := RawCanFrame{
		OID:  1,
		ID:   1,
		Rtr:  false,
		Eff:  false,
		Err:  false,
		Dlc:  1,
		Data: []byte{1},
	}
	result := ProcessedCanFrame{}
	ProcessRawCan(&result, testFrame)
	expected := "1#01"
	if result.PacketHash != expected {
		t.Errorf("%s != %s", result.PacketHash, expected)
	}
}

// BenchmarkProcessRawCan runs benchmarks on ProcessRawCan
func BenchmarkProcessRawCan(b *testing.B) {
	testFrame := RawCanFrame{
		OID:  1,
		ID:   1,
		Rtr:  false,
		Eff:  false,
		Err:  false,
		Dlc:  1,
		Data: []byte{1},
	}
	result := ProcessedCanFrame{}
	for i := 0; i < b.N; i++ {
		ProcessRawCan(&result, testFrame)
	}
}

// TestProcessCandump will verify that candump messages are appropriately converted to rawcanframe
func TestProcessCandump(t *testing.T) {
	expected := RawCanFrame{
		OID:       0,
		ID:        1,
		Rtr:       false,
		Eff:       false,
		Err:       false,
		Dlc:       1,
		Data:      []byte{1},
		Timestamp: 1000000000,
	}
	result := RawCanFrame{}
	ProcessCandump(&result, "(1) test 1#1")
	if result.OID != expected.OID || expected.ID != result.ID {
		t.Errorf("%v != %v", expected, result)
	}
}

// BenchmarkProcessCandump runs benchmarks on ProcessCandump
func BenchmarkProcessCandump(b *testing.B) {
	result := RawCanFrame{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ProcessCandump(&result, "(1) test 1#1")
	}
}

// TestProcessCanalyze will verify that canalyze dump messages are appropriately converted to rawcanframe
func TestProcessCanalyze(t *testing.T) {
	example := "vcan0   1513023350.178403       1     NOEFF   NORTR   NOERR   1     2       21 22"
	expected := RawCanFrame{
		OID:       1,
		ID:        1,
		Rtr:       false,
		Eff:       false,
		Err:       false,
		Dlc:       2,
		Data:      []byte{33, 34},
		Timestamp: 1513023350178402816,
	}
	result := RawCanFrame{}
	ProcessCanalyzeLog(&result, example)
	if result.OID != expected.OID || result.ID != expected.ID || result.Dlc != expected.Dlc {
		t.Errorf("%v != %v", expected, result)
	}
	if bytes.Compare(expected.Data, result.Data) != 0 {
		t.Errorf("%s != %s", expected.Data, result.Data)
	}

}
