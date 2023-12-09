package graphql

import (
	"github.com/Q00/go-chat/config"
	"github.com/Q00/go-chat/internal/domain/chat"
	"github.com/Q00/go-chat/internal/domain/room"
)

type Resolver struct {
	roomService *room.Service
	chatService *chat.Service
	pubsub      chat.PUBSUB
}

func NewResolver(cfg *config.AWSInfo) *Resolver {
	var p chat.PUBSUB
	p = chat.NewPUBSUB()
	go p.BroadCast()

	// initialize domain
	rs := room.NewService(cfg)
	cs := chat.NewService(cfg)

	return &Resolver{
		roomService: rs,
		chatService: cs,
		pubsub:      p,
	}
}
