package can

import (
	"errors"
	"net"

	"golang.org/x/sys/unix"
)

// InterfaceDescriptor is a can device handler that must be passed to functions using the can bus. Create one with SetupCanInterface.
type InterfaceDescriptor int

// SetupCanInterface will set up a CAN file descriptor to be used with sending and receiving CAN message.
// The function takes a string that specifies the interface to open.
func SetupCanInterface(canInterface string) (InterfaceDescriptor, error) {
	iface, err := net.InterfaceByName(canInterface)
	if err != nil {
		return 0, errors.New("error getting CAN interface by name: " + err.Error())
	}

	fd, err := unix.Socket(unix.AF_CAN, unix.SOCK_RAW, unix.CAN_RAW)
	if err != nil {
		return 0, errors.New("error setting CAN socket: " + err.Error())
	}

	addr := &unix.SockaddrCAN{Ifindex: iface.Index}

	err = unix.Bind(fd, addr)

	return InterfaceDescriptor(fd), err
}

// CloseCanInterface closes the passed interface
func CloseCanInterface(canInterface InterfaceDescriptor) error {
	err := unix.Close(int(canInterface))
	return err
}
