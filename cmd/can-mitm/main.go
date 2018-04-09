package main

import (
	"flag"
	"fmt"

	"github.com/jbreitbart/canlib"
)

type CanInstance struct {
	frame can.Frame
	src   int
}

func main() {
	canGlobalStr := flag.String("global", "", "The CAN interface for the greater CAN network")
	canTargetStr := flag.String("target", "", "The CAN interface for the targeted CAN device")
	flag.Parse()

	canGlobalIn := make(chan can.Frame, 100)
	canGlobalOut := make(chan can.Frame, 1)
	canTargetIn := make(chan can.Frame, 100)
	canTargetOut := make(chan can.Frame, 1)
	ctcInput := make(chan CanInstance)
	errChan := make(chan error)

	canGlobal, err := can.SetupCanInterface(*canGlobalStr)
	if err != nil {
		panic(err)
	}
	defer can.CloseCanInterface(canGlobal)

	canTarget, err := can.SetupCanInterface(*canTargetStr)
	if err != nil {
		panic(err)
	}
	defer can.CloseCanInterface(canTarget)

	go can.CaptureCan(canGlobal, canGlobalIn, errChan)
	go can.SendConcurrent(canGlobal, canGlobalOut, errChan)
	go processFrames(canGlobalIn, ctcInput, 1)

	go can.CaptureCan(canTarget, canTargetIn, errChan)
	go can.SendConcurrent(canTarget, canTargetOut, errChan)
	go processFrames(canTargetIn, ctcInput, 0)

	canTrafficControl(ctcInput, canTargetOut, canGlobalOut)
}

func canTrafficControl(input <-chan CanInstance, targetOut chan<- can.Frame, globalOut chan<- can.Frame) {
	history := []CanInstance{}
	printTemplate := "%s:\t%s\n"
	for update := range input {
		known := false
		var lastSeen CanInstance

		for _, entry := range history {
			known = can.CompareFrames(entry.frame, update.frame)
			if known != false {
				lastSeen = entry
				break
			}
		}

		if known == false {
			history = append(history, update)
			lastSeen = update
		}

		if update.src == 1 {
			fmt.Printf(printTemplate, "target", lastSeen.frame.ToString(" "))
			globalOut <- lastSeen.frame
		} else if update.src == 0 {
			fmt.Printf(printTemplate, "global", lastSeen.frame.ToString(" "))
			targetOut <- lastSeen.frame
		}
	}
}

func processFrames(captureChan <-chan can.Frame, ctcChan chan<- CanInstance, id int) {
	for newMessage := range captureChan {
		newInstance := CanInstance{
			frame: newMessage,
			src:   id,
		}
		ctcChan <- newInstance
	}
}
