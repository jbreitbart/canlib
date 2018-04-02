package canlib

import (
	"fmt"
	"testing"
)

// TestCreateRawFrame checks that CreateRawFrame appropriately creates a CanFrame
func TestCreateRawFrame(t *testing.T) {
	expected := CanFrame{Data: []byte{1}, OID: 1}
	result, err := CreateFrame(1, []byte{1}, false, false)
	if err != nil {
		t.Error("CreateRawFrame returned an error: " + err.Error())
	}
	if !CompareRawFrames(expected, result) {
		t.Errorf(fmt.Sprintf("%v != %v", result, expected))
	}
}
