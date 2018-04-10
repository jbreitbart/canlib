package can

import (
	"testing"
)

// TestFrameInSlice will check that FrameInSlice returns true if a given Frame exists in a slice
func TestFrameInSlice(t *testing.T) {
	testFrame := &Frame{OID: 1, Data: []byte{1}}
	testSlice := make([]*Frame, 0)
	testSlice = append(testSlice, testFrame)
	result := FrameInSlice(testFrame, testSlice)
	if result != true {
		t.Error("testFrame was not in testSlice")
	}
}

// TestFrameInSliceFail will check that FrameInSlice returns false if a given Frame does not exist in a slice
func TestFrameInSliceFail(t *testing.T) {
	testFrame := &Frame{OID: 1, Data: []byte{1}}
	testSlice := make([]*Frame, 0)
	result := FrameInSlice(testFrame, testSlice)
	if result == true {
		t.Error("testFrame was somehow in testSlice")
	}
}

// TestCompareFrames checks that CompareFrames returns true if two Frames are the same
func TestCompareFrames(t *testing.T) {
	testFrame := &Frame{OID: 1, Data: []byte{1}}
	result := CompareFrames(testFrame, testFrame)
	if result != true {
		t.Error("CompareFrames returned false when checking identical frames")
	}
}

// TestCompareFramesFail checks that CompareFrames returns false if mismatching Frames are compared
func TestCompareFramesFail(t *testing.T) {
	testFrameOne := &Frame{OID: 1, Data: []byte{1}}
	testFrameTwo := &Frame{OID: 2, Data: []byte{1}}
	result := CompareFrames(testFrameOne, testFrameTwo)
	if result == true {
		t.Error("CompareFrames returned true when checking mismatching frames")
	}
}

// TestCompareFramesSimple checks that CompareFramesSimple returns true if two Frames are the same
func TestCompareFramesSimple(t *testing.T) {
	testFrame := &Frame{OID: 1, Data: []byte{1}}
	result := CompareFramesSimple(testFrame, testFrame)
	if result != true {
		t.Error("CompareFramesSimple returned false when checking identical frames")
	}
}

// TestCompareFramesSimpleFail checks that CompareFrames returns false if mismatching Frames are compared
func TestCompareFramesSimpleFail(t *testing.T) {
	testFrameOne := &Frame{OID: 1, Data: []byte{1}}
	testFrameTwo := &Frame{OID: 2, Data: []byte{1}}
	result := CompareFramesSimple(testFrameOne, testFrameTwo)
	if result == true {
		t.Error("CompareFramesSimple returned true when checking mismatching frames")
	}
}

// TestFrameInSliceSimple will check that FrameInSliceSimple returns true if a given Frame exists in a slice
func TestFrameInSliceSimple(t *testing.T) {
	testFrame := &Frame{OID: 1, Data: []byte{1}}
	testSlice := make([]*Frame, 0)
	testSlice = append(testSlice, testFrame)
	result := FrameInSliceSimple(testFrame, testSlice)
	if result != true {
		t.Error("testFrame was not in testSlice")
	}
}

// TestFrameInSliceSimpleFail will check that FrameInSlice returns false if a given Frame does not exist in a slice
func TestFrameInSliceSimpleFail(t *testing.T) {
	testFrame := &Frame{OID: 1, Data: []byte{1}}
	testSlice := make([]*Frame, 0)
	result := FrameInSliceSimple(testFrame, testSlice)
	if result == true {
		t.Error("testFrame was somehow in testSlice")
	}
}
