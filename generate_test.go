package canlib

import (
	"fmt"
	"testing"
)

// TestCreateRawFrame checks that CreateRawFrame appropriately creates a RawCanFrame
func TestCreateRawFrame(t *testing.T) {
	expected := RawCanFrame{ID: 1, Eff: true, Rtr: false, Err: false, Dlc: 1, Data: []byte{1}, OID: 0}
	result, err := CreateRawFrame(1, []byte{1}, true, false, false)
	if err != nil {
		t.Error("CreateRawFrame returned an error: " + err.Error())
	}
	if !CompareRawFrames(expected, result) {
		t.Errorf(fmt.Sprintf("%v != %v", result, expected))
	}
}
