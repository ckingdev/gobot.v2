package bot

import (
	"github.com/cpalone/gobot.v2/proto"

	"euphoria.io/scope"
)

type ConnectionFactory func(ctx scope.Context, roomName string) proto.Connection
