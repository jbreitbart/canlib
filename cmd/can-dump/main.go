package main

import (
	"flag"
	"fmt"

	"github.com/jbreitbart/canlib"
)

func main() {
	caniface := flag.String("canif", "vcan0", "The CAN interface to capture on")
	flag.Parse()

	canFD, err := can.SetupCanInterface(*caniface)
	if err != nil {
		panic(err)
	}
	defer can.CloseCanInterface(canFD)

	c := make(chan *can.Frame, 100)
	errChan := make(chan error)
	go can.CaptureCan(canFD, c, errChan)
	go printCan(c)
	<-errChan
}

func printCan(ch <-chan *can.Frame) {
	for n := range ch {
		fmt.Println(n.ToString(" \t"))
	}
}
