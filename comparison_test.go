package canlib

import (
	"testing"
)

// TestRawFrameInSlice will check that RawFrameInSlice returns true if a given CanFrame exists in a slice
func TestRawFrameInSlice(t *testing.T) {
	testFrame := CanFrame{OID: 1, Data: []byte{1}}
	testSlice := make([]CanFrame, 1)
	testSlice = append(testSlice, testFrame)
	result := RawFrameInSlice(testFrame, testSlice)
	if result != true {
		t.Error("testFrame was not in testSlice")
	}
}

// TestRawFrameInSliceFail will check that RawFrameInSlice returns false if a given CanFrame does not exist in a slice
func TestRawFrameInSliceFail(t *testing.T) {
	testFrame := CanFrame{OID: 1, Data: []byte{1}}
	testSlice := make([]CanFrame, 1)
	result := RawFrameInSlice(testFrame, testSlice)
	if result == true {
		t.Error("testFrame was somehow in testSlice")
	}
}

// TestCompareRawFrames checks that CompareRawFrames returns true if two CanFrames are the same
func TestCompareRawFrames(t *testing.T) {
	testFrame := CanFrame{OID: 1, Data: []byte{1}}
	result := CompareRawFrames(testFrame, testFrame)
	if result != true {
		t.Error("CompareRawFrames returned false when checking identical frames")
	}
}

// TestCompareRawFramesFail checks that CompareRawFrames returns false if mismatching CanFrames are compared
func TestCompareRawFramesFail(t *testing.T) {
	testFrameOne := CanFrame{OID: 1, Data: []byte{1}}
	testFrameTwo := CanFrame{OID: 2, Data: []byte{1}}
	result := CompareRawFrames(testFrameOne, testFrameTwo)
	if result == true {
		t.Error("CompareRawFrames returned true when checking mismatching frames")
	}
}

// TestCompareRawFramesSimple checks that CompareRawFramesSimple returns true if two CanFrames are the same
func TestCompareRawFramesSimple(t *testing.T) {
	testFrame := CanFrame{OID: 1, Data: []byte{1}}
	result := CompareRawFramesSimple(testFrame, testFrame)
	if result != true {
		t.Error("CompareRawFramesSimple returned false when checking identical frames")
	}
}

// TestCompareRawFramesSimpleFail checks that CompareRawFrames returns false if mismatching CanFrames are compared
func TestCompareRawFramesSimpleFail(t *testing.T) {
	testFrameOne := CanFrame{OID: 1, Data: []byte{1}}
	testFrameTwo := CanFrame{OID: 2, Data: []byte{1}}
	result := CompareRawFramesSimple(testFrameOne, testFrameTwo)
	if result == true {
		t.Error("CompareRawFramesSimple returned true when checking mismatching frames")
	}
}

// TestRawFrameInSliceSimple will check that RawFrameInSliceSimple returns true if a given CanFrame exists in a slice
func TestRawFrameInSliceSimple(t *testing.T) {
	testFrame := CanFrame{OID: 1, Data: []byte{1}}
	testSlice := make([]CanFrame, 1)
	testSlice = append(testSlice, testFrame)
	result := RawFrameInSliceSimple(testFrame, testSlice)
	if result != true {
		t.Error("testFrame was not in testSlice")
	}
}

// TestRawFrameInSliceSimpleFail will check that RawFrameInSlice returns false if a given CanFrame does not exist in a slice
func TestRawFrameInSliceSimpleFail(t *testing.T) {
	testFrame := CanFrame{OID: 1, Data: []byte{1}}
	testSlice := make([]CanFrame, 1)
	result := RawFrameInSliceSimple(testFrame, testSlice)
	if result == true {
		t.Error("testFrame was somehow in testSlice")
	}
}
