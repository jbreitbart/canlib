package can

import (
	"fmt"
	"testing"

	"golang.org/x/sys/unix"
)

// TestCreateFrame checks that CreateFrame appropriately creates a Frame
func TestCreateFrame(t *testing.T) {
	expected := &Frame{Data: []byte{1}, OID: 1}
	result, err := CreateFrame(1, []byte{1}, false, false)
	if err != nil {
		t.Error("CreateFrame returned an error: " + err.Error())
	}
	if !CompareFrames(expected, result) {
		t.Errorf(fmt.Sprintf("%v != %v", result, expected))
	}

	result, err = CreateFrame(1, []byte{1}, true, false)
	if err != nil {
		t.Error("CreateFrame returned an error: " + err.Error())
	}
	if result.Rtr() != true {
		t.Error("CreateFrame failed to set Rtr.")
	}

	result, err = CreateFrame(1, []byte{1}, false, true)
	if err != nil {
		t.Error("CreateFrame returned an error: " + err.Error())
	}
	if result.Err() != true {
		t.Error("CreateFrame failed to set Err.")
	}

	result, err = CreateFrame(1, []byte{1}, true, true)
	if err != nil {
		t.Error("CreateFrame returned an error: " + err.Error())
	}
	if result.Err() != true {
		t.Error("CreateFrame failed to set Err.")
	}
	if result.Rtr() != true {
		t.Error("CreateFrame failed to set Rtr.")
	}
}

// TestBrokenCreateFrame checks that CreateFrame detects possible errors
func TestBrokenCreateFrame(t *testing.T) {
	_, err := CreateFrame(unix.CAN_EFF_MASK+1, []byte{1}, false, false)
	if err == nil {
		t.Error("CreateFrame created a frame with an invalid ID.")
	}

	_, err = CreateFrame(unix.CAN_EFF_MASK, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8}, false, false)
	if err == nil {
		t.Error("CreateFrame created a frame with more than 8 bytes of payload.")
	}
}
