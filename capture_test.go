package canlib

import (
	"fmt"
	"testing"

	"golang.org/x/sys/unix"
)

// TestCaptureSend Sends and captures a message
func TestCaptureSend(t *testing.T) {
	canFD0, err := SetupCanInterface("vcan0")
	if err != nil {
		t.Errorf("Failed can interface setup: %v", err)
	}

	canFD1, err := SetupCanInterface("vcan0")
	if err != nil {
		t.Errorf("Failed can interface setup: %v", err)
	}

	sFrame, err := CreateFrame(unix.CAN_EFF_MASK, []byte{1, 2, 3, 4}, false, false)
	if err != nil {
		t.Errorf("Failed creating frame: %v", err)
	}
	sFrame2, err := CreateFrame(unix.CAN_EFF_MASK, []byte{1, 2, 3, 5}, false, false)
	if err != nil {
		t.Errorf("Failed creating frame: %v", err)
	}

	equal, err := sendReceiveExpect(canFD0, canFD1, sFrame, sFrame)
	if err != nil {
		t.Errorf("Error while send/receive: %v", err)
	}
	if !equal {
		t.Error("Frames were not identical!")
	}

	equal, err = sendReceiveExpect(canFD0, canFD1, sFrame2, sFrame)
	if err != nil {
		t.Errorf("Error while send/receive: %v", err)
	}
	if equal {
		t.Error("Frames were identical, but should be different!")
	}

	err = CloseCanInterface(canFD0)
	if err != nil {
		t.Errorf("Failed can interface close: %v", err)
	}
	err = CloseCanInterface(canFD1)
	if err != nil {
		t.Errorf("Failed can interface close: %v", err)
	}

}

func sendReceiveExpect(canFD0 CANInterfaceDescriptor, canFD1 CANInterfaceDescriptor, sendFrame CanFrame, expectFrame CanFrame) (bool, error) {
	frameChan := make(chan CanFrame)
	errChan := make(chan error)

	go ReceiveNFrames(canFD0, 1, frameChan, errChan)
	err := SendCanFrame(canFD1, sendFrame)

	if err != nil {
		return false, fmt.Errorf("Error while sending frame: %v", err)
	}

	rFrame := <-frameChan
	return CompareRawFrames(rFrame, expectFrame), nil
}
