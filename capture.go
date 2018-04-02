package canlib

import (
	"fmt"
	"time"

	"golang.org/x/sys/unix"
)

// ReceiveFrame reads a frame from the provided interface
func ReceiveFrame(canInterface CANInterfaceDescriptor) (CanFrame, error) {
	var frame CanFrame
	bytes := make([]byte, 16)
	n, err := unix.Read(int(canInterface), bytes)
	if err != nil {
		return frame, err
	}
	if n != 16 {
		return frame, fmt.Errorf("Could only write %v of 16 bytes to socket.", n)
	}

	captime := time.Now().UnixNano()
	frame, err = CreateFrameFromByte(bytes, captime)
	if err != nil {
		return frame, err
	}

	return frame, nil
}

// ReceiveNFrames reads N frames from the provided interface
func ReceiveNFrames(canInterface CANInterfaceDescriptor, n int, canChannel chan<- CanFrame, errorChannel chan<- error) {
	for i := 0; i < n; i++ {
		frame, err := ReceiveFrame(canInterface)
		if err != nil {
			errorChannel <- err
		}

		canChannel <- frame
	}
}

// CaptureCan will listen to the provided SocketCAN interface and add any messages seen to the provided channel
// TODO remove?
func CaptureCan(canInterface CANInterfaceDescriptor, canChannel chan<- CanFrame, errorChannel chan<- error) {
	for {
		ReceiveNFrames(canInterface, 1000, canChannel, errorChannel)
	}
}
