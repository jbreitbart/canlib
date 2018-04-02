package main

import (
	"flag"
	"fmt"

	"github.com/jbreitbart/canlib"
)

func main() {
	caniface := flag.String("canif", "vcan0", "The CAN interface to capture on")
	flag.Parse()

	canFD, err := canlib.SetupCanInterface(*caniface)
	if err != nil {
		panic(err)
	}
	defer canlib.CloseCanInterface(canFD)

	c := make(chan canlib.CanFrame, 100)
	errChan := make(chan error)
	go canlib.CaptureCan(canFD, c, errChan)
	go printCan(c)
	<-errChan
}

func printCan(ch <-chan canlib.CanFrame) {
	for n := range ch {
		fmt.Println(canlib.CanFrameToString(n, " \t"))
	}
}
