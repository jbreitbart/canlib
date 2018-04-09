package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/jbreitbart/canlib"
)

func main() {
	canifaceIn := flag.String("global", "vcan0", "The CAN interface to capture on")
	canifaceOut := flag.String("target", "vcan1", "The CAN interface to pipe to")
	flag.Parse()

	canFDIn, err := can.SetupCanInterface(*canifaceIn)
	if err != nil {
		panic(err)
	}
	defer can.CloseCanInterface(canFDIn)

	canFDOut, err := can.SetupCanInterface(*canifaceOut)
	if err != nil {
		panic(err)
	}
	defer can.CloseCanInterface(canFDOut)

	canGlobalChan := make(chan can.Frame, 100)
	canTargetChan := make(chan can.Frame, 100)
	canMultiplexOne := make(chan can.Frame, 100)
	canMultiplexTwo := make(chan can.Frame, 100)
	output := make(chan can.Frame, 100)
	errChan := make(chan error)

	go can.CaptureCan(canFDIn, canGlobalChan, errChan)
	go can.CaptureCan(canFDOut, canTargetChan, errChan)
	go can.SendConcurrent(canFDOut, canMultiplexOne, errChan)
	go globalMultiplex(canGlobalChan, canMultiplexOne, canMultiplexTwo)
	go processing(canTargetChan, canMultiplexTwo, output)

	for message := range output {
		fmt.Println(message)
	}

}

// globalMultiplex will read a value from globalChan and sent that value to both mplexOne and mplexTwo
func globalMultiplex(globalChan <-chan can.Frame, mplexOne chan<- can.Frame,
	mplexTwo chan<- can.Frame) {

	for message := range globalChan {
		mplexOne <- message
		mplexTwo <- message

	}
}

// processing will start another process to load an array of known messages and then diff that with the target captures
func processing(targetChan <-chan can.Frame, globalChan <-chan can.Frame, output chan<- can.Frame) {
	var seenMessages = []can.Frame{}
	var mutex = &sync.Mutex{}

	go func() {
		for globalMessage := range globalChan {
			mutex.Lock()
			if can.FrameInSlice(globalMessage, seenMessages) == false {
				seenMessages = append(seenMessages, globalMessage)
			}
			mutex.Unlock()
		}
	}()

	for newMessage := range targetChan {
		mutex.Lock()
		if can.FrameInSlice(newMessage, seenMessages) == false {
			newMessage.ToString("\t")
		}
		mutex.Unlock()
	}
}
