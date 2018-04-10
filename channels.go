package can

// ChannelMultiplex will take a Frame sent into the input channel and relay it to all output channels
func ChannelMultiplex(input <-chan *Frame, output ...chan<- *Frame) {

	for message := range input {
		for _, out := range output {
			out <- message
		}
	}
}
