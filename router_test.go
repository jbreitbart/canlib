package can

import (
	"fmt"
	"testing"

	"golang.org/x/sys/unix"
)

// TestRouter tests the basic router functionality
func TestRouter(t *testing.T) {
	canFD0, err := SetupCanInterface("vcan0")
	if err != nil {
		t.Errorf("Failed can interface setup: %v", err)
	}
	canFD1, err := SetupCanInterface("vcan0")
	if err != nil {
		t.Errorf("Failed can interface setup: %v", err)
	}

	r := NewRouter(canFD0)

	ch := make(chan *Frame, 2)
	r.Subscribe(100, ch)

	equal, err := sendReceiveExpectRouterFilter(canFD1, ch)
	if err != nil {
		t.Errorf("Error while send/receive router: %v", err)
	}
	if !equal {
		t.Error("Frames were not identical!")
	}

	r.Unsubscribe(100, ch)
	r.Subscribe(ALLIDs, ch)

	equal, err = sendReceiveExpectRouterAllIDs(canFD1, ch)
	if err != nil {
		t.Errorf("Error while send/receive router: %v", err)
	}
	if !equal {
		t.Error("Frames were not identical!")
	}

	r.Stop()

	err = CloseCanInterface(canFD0)
	if err != nil {
		t.Errorf("Failed can interface close: %v", err)
	}
	err = CloseCanInterface(canFD1)
	if err != nil {
		t.Errorf("Failed can interface close: %v", err)
	}
}

func sendReceiveExpectRouterFilter(canFD InterfaceDescriptor, expect chan *Frame) (bool, error) {
	ignoredFrame, err := CreateFrame(unix.CAN_EFF_MASK, []byte{1, 2, 3, 5}, false, false)
	if err != nil {
		return false, fmt.Errorf("Failed creating frame: %v", err)
	}

	err = SendFrame(canFD, ignoredFrame)
	if err != nil {
		return false, fmt.Errorf("Error while sending frame: %v", err)
	}

	sFrame, err := CreateFrame(100, []byte{1, 2, 3, 4}, false, false)
	if err != nil {
		return false, fmt.Errorf("Failed creating frame: %v", err)
	}

	err = SendFrame(canFD, sFrame)
	if err != nil {
		return false, fmt.Errorf("Error while sending frame: %v", err)
	}

	rFrame := <-expect
	return CompareFrames(rFrame, sFrame), nil
}

func sendReceiveExpectRouterAllIDs(canFD InterfaceDescriptor, expect chan *Frame) (bool, error) {
	sFrame0, err := CreateFrame(unix.CAN_EFF_MASK, []byte{1, 2, 3, 5}, false, false)
	if err != nil {
		return false, fmt.Errorf("Failed creating frame: %v", err)
	}

	err = SendFrame(canFD, sFrame0)
	if err != nil {
		return false, fmt.Errorf("Error while sending frame: %v", err)
	}

	sFrame1, err := CreateFrame(100, []byte{1, 2, 3, 4}, false, false)
	if err != nil {
		return false, fmt.Errorf("Failed creating frame: %v", err)
	}

	err = SendFrame(canFD, sFrame1)
	if err != nil {
		return false, fmt.Errorf("Error while sending frame: %v", err)
	}

	rFrame := <-expect
	equal := CompareFrames(rFrame, sFrame0)

	rFrame = <-expect
	equal = equal && CompareFrames(rFrame, sFrame1)

	return equal, nil
}
