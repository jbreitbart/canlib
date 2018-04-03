package canlib

import (
	"bytes"
)

// CompareRawFrames takes two Raw Can Frames and returns true if they are the same frame and false otherwise
//
// This comparison is done on all fields and flags except anything time based.
// Since a Raw Can Frame's OID contains the masked ID and Flags, it is used for comparison to save a bit of computation.
// Because of this OID comparison, this function is not compatible with CanFrame structs that are built with
// SocketCan's candump output is not supported. Instead use CompareRawFramesSimple instead.
func CompareRawFrames(frameOne CanFrame, frameTwo CanFrame) bool {
	if (frameOne.OID == frameTwo.OID) && (frameOne.Dlc() == frameTwo.Dlc()) {
		if bytes.Equal(frameOne.Data, frameTwo.Data) {
			return true
		}
	}
	return false
}

// CompareRawFramesSimple takes two CanFrames and returns true if they are the same frame and false otherwise
//
// This comparison is only performed on the ID, Data Length, and Data Contents. It does not support checking flasgs
// or masks in order to support CanFrames that are built from SocketCan's candump output.
func CompareRawFramesSimple(frameOne CanFrame, frameTwo CanFrame) bool {
	if (frameOne.ID() == frameTwo.ID()) && (frameOne.Dlc() == frameTwo.Dlc()) {
		if bytes.Equal(frameOne.Data, frameTwo.Data) {
			return true
		}
	}
	return false
}

// RawFrameInSlice takes a Raw Can Frame and looks to see if it exists within a slice of Raw Can Frames
//
// Because this function makes use of CompareRawFrames, it is not compatible with CanFrames that are
// built from SocketCan's candump output. Instead, use RawFrameInSliceSimple.
func RawFrameInSlice(frame CanFrame, frameSlice []CanFrame) bool {
	for _, slice := range frameSlice {
		if CompareRawFrames(frame, slice) {
			return true
		}
	}
	return false
}

// RawFrameInSliceSimple takes a CanFrame and looks to see if it exists within a slice of CanFrames using the simple method
//
// Because this function makes use of CompareRawFramesSimple, it is compatible with CanFrames that are built
// from SocketCan's candump output.
func RawFrameInSliceSimple(frame CanFrame, frameSlice []CanFrame) bool {
	for _, slice := range frameSlice {
		if CompareRawFramesSimple(frame, slice) {
			return true
		}
	}
	return false
}
