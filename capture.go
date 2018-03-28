package canlib

import (
	"fmt"
	"time"

	"golang.org/x/sys/unix"
)

// CaptureCan will listen to the provided SocketCAN interface and add any messages seen to the provided channel
func CaptureCan(canInterface CANInterfaceDescriptor, canChannel chan<- RawCanFrame, errorChannel chan<- error) {
	frame := make([]byte, 16)
	canmsg := new(RawCanFrame)
	for {
		n, err := unix.Read(int(canInterface), frame)
		if err != nil {
			errorChannel <- err
		}
		if n != 16 {
			errorChannel <- fmt.Errorf("Could only write %v of 16 bytes to socket.", n)
		}

		captime := time.Now().UnixNano()
		ByteArrayToCanFrame(frame, canmsg, captime)
		canChannel <- *canmsg
	}
}
