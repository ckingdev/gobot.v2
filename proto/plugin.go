package proto

import (
	hproto "euphoria.io/heim/proto"
)

// Plugin describes the smallest unit of functionality. A Room may have multiple
// Plugins, but a Plugin has exactly one Room. A Plugin consumes packets from a
// Room's Connection and sends any replies through the Room.
type Plugin interface {
	// Room returns the Room that the Plugin is associated with.
	Room() Room

	// Start begins the Plugin's main functionality. The only design requirement
	// is that the plugin must consume Packets promptly even if it does not use
	// those Packets, so that the Room's packet dispatch loop does not block.
	Start() error

	// Kill attempts to cleanly shut down the Plugin and all associated goroutines.
	Kill() error

	// Incoming returns a send-only channel which consumes incoming packets from
	// the associated Room's Connection and responds accordingly.
	Incoming() chan<- *hproto.Packet
}
