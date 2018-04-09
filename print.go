package can

import (
	"fmt"
	"strconv"
)

// ToString takes a Frame and makes it look pretty based on several parameters
//
// This function is designed to be used to prepare a Frame for multiple output formats
// including stdout, csv, and other custom delimited formats.
func (frame Frame) ToString(delimiter string) string {
	var frameString string
	timestamp := TimestampToSeconds(frame.Timestamp)
	frameString += strconv.FormatFloat(timestamp, 'f', -1, 64) + delimiter
	frameString += fmt.Sprintf("%X", frame.OID) + delimiter
	if frame.Eff() {
		frameString += "EFF" + delimiter
	} else {
		frameString += "NOEFF" + delimiter
	}
	if frame.Rtr() {
		frameString += "RTR" + delimiter
	} else {
		frameString += "NORTR" + delimiter
	}
	if frame.Err() {
		frameString += "ERR" + delimiter
	} else {
		frameString += "NOERR" + delimiter
	}
	frameString += fmt.Sprintf("%X", frame.ID()) + delimiter
	frameString += fmt.Sprintf("%d", frame.Dlc()) + delimiter
	frameString += fmt.Sprintf("% X", frame.Data)
	return frameString
}

// TimestampToSeconds converts the int64 timestamp into a float unix timestamp
func TimestampToSeconds(timestamp int64) float64 {
	floatTime := float64(timestamp)
	return floatTime * 0.000000001
}

// ToString takes a ProcessedFrame and formats it based on several parameters
//
// This function is designed to be used to prepare a Frame for multiple output formats
// including stdout, csv, and other custom delimited formats.
func (frame ProcessedFrame) ToString(delimiter string) string {
	var frameString string
	frameString += frame.PacketHash + delimiter
	frameString += frame.Packet.ToString(delimiter)
	return frameString
}
