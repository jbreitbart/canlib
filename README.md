# CANLIB - A GO library and a series of utilities for CAN bus testing

## Install

-   Install libraries and utilities: `> go get github.com/buffersandbeer/canlib/...`
-   Install just the library: `> go get github.com/buffersandbeer/canlib/`

## Userspace Utilities

-   `can-dump` - Dump CAN packets from SocketCan interface and display extended information
-   `can-fuzz` - Incrementally fuzz CAN messages
-   `can-halfpipe` - Print messages originiating from a target device using a "bump in the wire"

## Docs

Documentation and usage explanations for the library can be found at <https://godoc.org/github.com/buffersandbeer/canlib>.

## Tests

`go test` is used for unit testing. Tests require a vcan interface to be successful:

    sudo modprobe vcan
    sudo ip link add dev vcan0 type vcan
    sudo ip link set up vcan0


## Library Features

-   Write to CAN Bus interface
-   Read from CAN Bus interface
-   Generate CAN messages
-   Process CAN messages
-   Pretty Print CAN messages
