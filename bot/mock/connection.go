package mock

import (
	"github.com/cpalone/gobot.v2/proto"

	hproto "euphoria.io/heim/proto"
	"euphoria.io/scope"
)

type Connection struct {
	ctx scope.Context

	testToClient chan *hproto.Packet
	clientToTest chan *hproto.Packet

	incoming chan *hproto.Packet
	outgoing chan *hproto.Packet
}

func (c *Connection) Connect() error {
	return nil
}

func (c *Connection) Run() error {
	for {
		select {
		case msg := <-c.outgoing:
			c.clientToTest <- msg
		case msg := <-c.testToClient:
			c.incoming <- msg
		case <-c.ctx.Done():
			return nil
		}
	}
}

func (c *Connection) Kill() error {
	c.ctx.Cancel()
	return nil
}

func (c *Connection) Incoming() <-chan *hproto.Packet {
	return c.incoming
}

func (c *Connection) Outgoing() chan<- *hproto.Packet {
	return c.outgoing
}

func WSConnectionFactory(ctx scope.Context, roomName string) proto.Connection {
	return &Connection{
		ctx: ctx.Fork(),

		testToClient: make(chan *hproto.Packet),
		clientToTest: make(chan *hproto.Packet),

		incoming: make(chan *hproto.Packet),
		outgoing: make(chan *hproto.Packet),
	}
}
