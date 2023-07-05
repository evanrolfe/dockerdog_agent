package parse

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
)

// httpStreamFactory implements tcpassembly.StreamFactory
type HttpStreamFactory struct{}

// httpStream will handle the actual decoding of http requests.
type HttpStream struct {
	net, transport gopacket.Flow
	reader         tcpreader.ReaderStream
}

func (h *HttpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hstream := &HttpStream{
		net:       net,
		transport: transport,
		reader:    tcpreader.NewReaderStream(),
	}
	go hstream.run() // Important... we must guarantee that data from the reader stream is read.

	// ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
	return &hstream.reader
}

func (h *HttpStream) run() {
	fmt.Println("Net:", h.net.String())
	fmt.Println("Transport:", h.transport.String())

	buffer := make([]byte, 1024) // Adjust the buffer size as needed]
	for {
		n, err := h.reader.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from stream:", err)
			break
		}

		// Process the read bytes
		data := buffer[:n]
		// fmt.Println("Received", n, "bytes:", data)
		fmt.Println(string(data))
		// Check if the end of the stream has been reached
		if n == 0 {
			fmt.Println("End of stream")
			break
		}
	}
}