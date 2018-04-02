package canlib

import "testing"

// SetupCANInterface opens and closes a vCAN interface
func TestSetupCANInterface(t *testing.T) {

	canFD, err := SetupCanInterface("vcan0")
	if err != nil {
		t.Errorf("Failed can interface setup: %v", err)
	}

	err = CloseCanInterface(canFD)
	if err != nil {
		t.Errorf("Failed can interface close: %v", err)
	}

	canFD, err = SetupCanInterface("asdf")
	if err == nil {
		t.Error("Could open CAN device that does not exist: asdf")
	}

}
