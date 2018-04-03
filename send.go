package canlib

import (
	"encoding/binary"
	"errors"
	"fmt"

	"golang.org/x/sys/unix"
)

// SendCanFrame will send the provided CAN message on the given CAN interface
func SendCanFrame(canInterface CANInterfaceDescriptor, message CanFrame) error {
	if message.Dlc() > 8 {
		return errors.New("CAN message to send is invalid")
	}

	frame := make([]byte, 16)
	binary.LittleEndian.PutUint32(frame[0:4], message.OID)
	frame[4] = byte(message.Dlc())
	copy(frame[8:], message.Data)
	n, err := unix.Write(int(canInterface), frame)
	if n != 16 {
		return fmt.Errorf("Only %v of 16 bytes written to socket", n)
	}

	return err
}

// SendCanConcurrent will utilize a channel to send CAN messages on the given CAN interface
func SendCanConcurrent(canInterface CANInterfaceDescriptor, canChannel <-chan CanFrame, errorChannel chan<- error) {
	for message := range canChannel {
		err := SendCanFrame(canInterface, message)
		if err != nil {
			errorChannel <- err
		}
	}
}
