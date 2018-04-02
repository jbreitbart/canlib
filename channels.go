package canlib

// RawCanChannelMultiplex will take a CanFrame sent into the input channel and relay it to all output channels
func RawCanChannelMultiplex(input <-chan CanFrame, output ...chan<- CanFrame) {

	for message := range input {
		for _, out := range output {
			out <- message
		}
	}
}
