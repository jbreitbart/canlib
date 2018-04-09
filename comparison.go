package can

import (
	"bytes"
)

// CompareFrames takes two Raw Can Frames and returns true if they are the same frame and false otherwise
//
// This comparison is done on all fields and flags except anything time based.
// Since a Raw Can Frame's OID contains the masked ID and Flags, it is used for comparison to save a bit of computation.
// Because of this OID comparison, this function is not compatible with Frame structs that are built with
// SocketCan's candump output is not supported. Instead use CompareFramesSimple instead.
func CompareFrames(frameOne Frame, frameTwo Frame) bool {
	if (frameOne.OID == frameTwo.OID) && (frameOne.Dlc() == frameTwo.Dlc()) {
		if bytes.Equal(frameOne.Data, frameTwo.Data) {
			return true
		}
	}
	return false
}

// CompareFramesSimple takes two Frames and returns true if they are the same frame and false otherwise
//
// This comparison is only performed on the ID, Data Length, and Data Contents. It does not support checking flasgs
// or masks in order to support Frames that are built from SocketCan's candump output.
func CompareFramesSimple(frameOne Frame, frameTwo Frame) bool {
	if (frameOne.ID() == frameTwo.ID()) && (frameOne.Dlc() == frameTwo.Dlc()) {
		if bytes.Equal(frameOne.Data, frameTwo.Data) {
			return true
		}
	}
	return false
}

// FrameInSlice takes a Raw Can Frame and looks to see if it exists within a slice of Raw Can Frames
//
// Because this function makes use of CompareFrames, it is not compatible with Frames that are
// built from SocketCan's candump output. Instead, use FrameInSliceSimple.
func FrameInSlice(frame Frame, frameSlice []Frame) bool {
	for _, slice := range frameSlice {
		if CompareFrames(frame, slice) {
			return true
		}
	}
	return false
}

// FrameInSliceSimple takes a Frame and looks to see if it exists within a slice of Frames using the simple method
//
// Because this function makes use of CompareFramesSimple, it is compatible with Frames that are built
// from SocketCan's candump output.
func FrameInSliceSimple(frame Frame, frameSlice []Frame) bool {
	for _, slice := range frameSlice {
		if CompareFramesSimple(frame, slice) {
			return true
		}
	}
	return false
}
