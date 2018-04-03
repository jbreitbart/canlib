package canlib

import (
	"encoding/binary"
	"errors"

	"golang.org/x/sys/unix"
)

// CreateFrameFromByte will take an byte array and tries to create a can CanFrame
func CreateFrameFromByte(array []byte, captureTime int64) (CanFrame, error) {
	var ret CanFrame

	ret.OID = binary.LittleEndian.Uint32(array[0:4])

	Dlc := array[4]
	if Dlc > 8 {
		return ret, errors.New("data too long. Data must be < 8 bytes")
	}
	ret.Data = array[8 : 8+Dlc]

	ret.Timestamp = captureTime

	return ret, nil
}

// CreateFrame will take an ID, Data, and Flags to generate a valid CanFrame
func CreateFrame(id uint32, data []byte, rtr bool, err bool) (CanFrame, error) {
	var ret CanFrame

	if id > unix.CAN_EFF_MASK {
		return ret, errors.New("ID too large.")
	}

	ret.OID = id

	// use extened id?
	if ret.OID > unix.CAN_SFF_MASK {
		ret.OID = ret.OID | unix.CAN_EFF_FLAG
	}

	if rtr {
		ret.OID = ret.OID | unix.CAN_RTR_FLAG
	}

	if err {
		ret.OID = ret.OID | unix.CAN_ERR_FLAG
	}

	dlc := uint8(len(data))
	if dlc > 8 {
		return ret, errors.New("data too long. Data must be < 8 bytes")
	}
	ret.Data = data

	return ret, nil
}
