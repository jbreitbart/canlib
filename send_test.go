package can

import (
	"testing"
)

func TestSendToLarge(t *testing.T) {
	canFD, err := SetupCanInterface("vcan0")
	if err != nil {
		t.Errorf("Failed can interface setup: %v", err)
	}

	sFrame, err := CreateFrame(100, []byte{1, 2, 3, 4}, false, false)
	if err != nil {
		t.Errorf("Failed creating frame: %v", err)
	}

	sFrame.Data = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	err = SendFrame(canFD, sFrame)
	if err == nil {
		t.Error("Sent a frame with more than 8 bytes of payload.")
	}

	err = CloseCanInterface(canFD)
	if err != nil {
		t.Errorf("Failed can interface close: %v", err)
	}
}
