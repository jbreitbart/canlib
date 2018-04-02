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

	canFDIn, err := canlib.SetupCanInterface(*canifaceIn)
	if err != nil {
		panic(err)
	}
	defer canlib.CloseCanInterface(canFDIn)

	canFDOut, err := canlib.SetupCanInterface(*canifaceOut)
	if err != nil {
		panic(err)
	}
	defer canlib.CloseCanInterface(canFDOut)

	canGlobalChan := make(chan canlib.CanFrame, 100)
	canTargetChan := make(chan canlib.CanFrame, 100)
	canMultiplexOne := make(chan canlib.CanFrame, 100)
	canMultiplexTwo := make(chan canlib.CanFrame, 100)
	output := make(chan canlib.CanFrame, 100)
	errChan := make(chan error)

	go canlib.CaptureCan(canFDIn, canGlobalChan, errChan)
	go canlib.CaptureCan(canFDOut, canTargetChan, errChan)
	go canlib.SendCanConcurrent(canFDOut, canMultiplexOne, errChan)
	go globalMultiplex(canGlobalChan, canMultiplexOne, canMultiplexTwo)
	go processing(canTargetChan, canMultiplexTwo, output)

	for message := range output {
		fmt.Println(message)
	}

}

// globalMultiplex will read a value from globalChan and sent that value to both mplexOne and mplexTwo
func globalMultiplex(globalChan <-chan canlib.CanFrame, mplexOne chan<- canlib.CanFrame,
	mplexTwo chan<- canlib.CanFrame) {

	for message := range globalChan {
		mplexOne <- message
		mplexTwo <- message

	}
}

// processing will start another process to load an array of known messages and then diff that with the target captures
func processing(targetChan <-chan canlib.CanFrame, globalChan <-chan canlib.CanFrame, output chan<- canlib.CanFrame) {
	var seenMessages = []canlib.CanFrame{}
	var mutex = &sync.Mutex{}

	go func() {
		for globalMessage := range globalChan {
			mutex.Lock()
			if canlib.RawFrameInSlice(globalMessage, seenMessages) == false {
				seenMessages = append(seenMessages, globalMessage)
			}
			mutex.Unlock()
		}
	}()

	for newMessage := range targetChan {
		mutex.Lock()
		if canlib.RawFrameInSlice(newMessage, seenMessages) == false {
			canlib.CanFrameToString(newMessage, "\t")
		}
		mutex.Unlock()
	}
}
