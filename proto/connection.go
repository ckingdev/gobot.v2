package proto

import (
	hproto "euphoria.io/heim/proto"
)

// Connection describes a connection (websocket or mock) to a heim server. A
// Connection is attached to exactly one Room.
type Connection interface {
	// Connect initiates the connection to the heim server.
	Connect() error

	// Run begins listening for packets coming from the associated Room and packets
	// received from the heim server.
	Run() error

	// Kill attempts to gracefully shut down the Connection and any associated
	// goroutines
	Kill() error

	// Incoming returns a channel that will have Packets coming from the server
	// placed as they are received.
	Incoming() <-chan *hproto.Packet

	// Outgoing returns a channel that Run() consumes from and sends over the
	// connection to the heim server.
	Outgoing() chan<- *hproto.Packet
}
