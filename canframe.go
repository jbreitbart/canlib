package canlib

import "golang.org/x/sys/unix"

// CanFrame represents the data contained in a CAN packet
type CanFrame struct {
	OID       uint32 // 32-bit CAN_ID including masks
	Data      []byte // Message Payload
	Timestamp int64  // Time message was captured as Unix Timestamp in nanoseconds
}

func (frame CanFrame) ID() uint32 {
	// remove the top 3 bits
	return frame.OID & unix.CAN_EFF_MASK
}

func (frame CanFrame) Rtr() bool {
	return (frame.OID & unix.CAN_RTR_FLAG) != 0
}

func (frame CanFrame) Err() bool {
	return (frame.OID & unix.CAN_ERR_FLAG) != 0
}

func (frame CanFrame) Eff() bool {
	return (frame.OID & unix.CAN_EFF_FLAG) != 0
}

func (frame CanFrame) Dlc() int {
	return len(frame.Data)
}

// ProcessedCanFrame represents a CAN packet and additional data about the packet
type ProcessedCanFrame struct {
	Packet       CanFrame // CAN packet
	PacketHash   string   // md5 hash of the Packet's ID and Data fields
	AlphaNumData string   // Any Alpha-numeric data within the can payload
}
