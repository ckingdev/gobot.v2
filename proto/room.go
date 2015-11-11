package proto

import (
	hproto "euphoria.io/heim/proto"
)

// Room describes the presence of a bot in a single heim room. A Room has exactly
// one Connection. It may have many Plugins. The Room is responsible for
// communication between Plugins and the Connection.
type Room interface {
	// AddPlugin attaches a Plugin to the Room and registers the plugin to be
	// forwarded incoming packets from the room.
	AddPlugin(plg Plugin) error

	// Start begins listening for packets on the attached Connection, forwarding
	// them to plugins, and forwarding packets from the plugins to the Connection.
	Start() error

	// Kill shuts down the room, its Connection, and any associated Plugins.
	Kill() error

	// SendPayload forms a packet from the given payload and type and sends it over
	// the Room's Connection.
	SendPayload(payload interface{}, pType hproto.PacketType) (*hproto.Packet, error)

	// SetConnection sets the Room to use the given Connection for all of its
	// communication with the heim server.
	SetConnection(c Connection)

	// GetConnection returns the Room's underlying Connection to the heim server.
	GetConnection() Connection

	// Name returns the name of the room on a heim server that the Room is attached
	// to.
	Name() string
}
