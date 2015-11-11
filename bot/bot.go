package bot

import (
	"github.com/cpalone/gobot.v2/proto"

	"euphoria.io/scope"
)

type Bot struct {
	rooms map[string]proto.Room
	ctx   scope.Context
}

func NewBot(ctx scope.Context) *Bot {
	return &Bot{
		ctx:   ctx.Fork(),
		rooms: make(map[string]proto.Room),
	}
}

func (b *Bot) AddRoom(r proto.Room) {
	b.rooms[r.Name()] = r
}
