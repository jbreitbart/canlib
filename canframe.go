package canlib

import "golang.org/x/sys/unix"

// CanFrame represents the data contained in a CAN packet
type CanFrame struct {
	OID       uint32 // 32-bit CAN_ID including masks
	Data      []byte // Message Payload
	Timestamp int64  // Time message was captured as Unix Timestamp in nanoseconds
}

// ID returns the ID of the frame
func (frame CanFrame) ID() uint32 {
	// remove the top 3 bits
	return frame.OID & unix.CAN_EFF_MASK
}

// Rtr returns the remote transmission request bit is set
func (frame CanFrame) Rtr() bool {
	return (frame.OID & unix.CAN_RTR_FLAG) != 0
}

// Err returns if the frame is an error frame
func (frame CanFrame) Err() bool {
	return (frame.OID & unix.CAN_ERR_FLAG) != 0
}

// Eff returns if this frame is using extended IDs (29bit)
func (frame CanFrame) Eff() bool {
	return (frame.OID & unix.CAN_EFF_FLAG) != 0
}

// Dlc returns the length of the data stored in the frame
func (frame CanFrame) Dlc() int {
	return len(frame.Data)
}

// ProcessedCanFrame represents a CAN packet and additional data about the packet
type ProcessedCanFrame struct {
	Packet       CanFrame // CAN packet
	PacketHash   string   // md5 hash of the Packet's ID and Data fields
	AlphaNumData string   // Any Alpha-numeric data within the can payload
}
