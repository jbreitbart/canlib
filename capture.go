package canlib

import (
	"fmt"
	"time"

	"golang.org/x/sys/unix"
)

// CaptureCan will listen to the provided SocketCAN interface and add any messages seen to the provided channel
func CaptureCan(canInterface CANInterfaceDescriptor, canChannel chan<- CanFrame, errorChannel chan<- error) {
	bytes := make([]byte, 16)
	for {
		n, err := unix.Read(int(canInterface), bytes)
		if err != nil {
			errorChannel <- err
		}
		if n != 16 {
			errorChannel <- fmt.Errorf("Could only write %v of 16 bytes to socket.", n)
		}

		captime := time.Now().UnixNano()
		frame, err := CreateFrameFromByte(bytes, captime)
		if err != nil {
			errorChannel <- err
		}

		canChannel <- frame
	}
}
