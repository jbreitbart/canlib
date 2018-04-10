package can

import (
	"encoding/binary"
	"errors"
	"fmt"

	"golang.org/x/sys/unix"
)

// SendFrame will send the provided CAN message on the given CAN interface
func SendFrame(canInterface InterfaceDescriptor, message *Frame) error {
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

// SendConcurrent will utilize a channel to send CAN messages on the given CAN interface
func SendConcurrent(canInterface InterfaceDescriptor, canChannel <-chan *Frame, errorChannel chan<- error) {
	for message := range canChannel {
		err := SendFrame(canInterface, message)
		if err != nil {
			errorChannel <- err
		}
	}
}
