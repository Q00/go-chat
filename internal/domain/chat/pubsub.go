package chat

import (
	"fmt"
	"github.com/Q00/go-chat/internal/dto"
	"github.com/Q00/go-chat/util"
	"sync"
	"time"
)

var (
	subscribers map[string]*chatSubscription
	mu          sync.Mutex
)

type PUBSUB interface {
	Subscribe(roomID string) (string, <-chan *dto.ChatMessage)
	UnSubscribe(ID string)
	BroadCast()
	Publish(message *dto.ChatMessage)
}

type chatSubscription struct {
	ID     string
	stop   chan struct{}
	roomID string
	events chan<- *dto.ChatMessage
}

type pubsub struct {
	chatAdded        chan *dto.ChatMessage
	chatSubscription chan *chatSubscription
}

func NewPUBSUB() *pubsub {
	p := &pubsub{
		chatAdded:        make(chan *dto.ChatMessage),
		chatSubscription: make(chan *chatSubscription),
	}

	return p
}
func (p *pubsub) Subscribe(roomID string) (string, <-chan *dto.ChatMessage) {
	events := make(chan *dto.ChatMessage)
	stop := make(chan struct{})
	UUID, _ := util.CreateUUID()
	ID := UUID.String()

	p.chatSubscription <- &chatSubscription{
		ID:     ID,
		stop:   stop,
		roomID: roomID,
		events: events,
	}

	fmt.Println("Subscribed to chat", ID)

	return ID, events
}

func (p *pubsub) UnSubscribe(ID string) {
	mu.Lock()
	defer mu.Unlock()

	if sub, ok := subscribers[ID]; ok {
		fmt.Println("Unsubscribed from chat", ID)
		close(sub.stop)
		delete(subscribers, ID)
	}
}

func (p *pubsub) BroadCast() {
	subscribers = map[string]*chatSubscription{}
	unsubscribe := make(chan string)

	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
			fmt.Println("Unsubscribed from chat", subscribers)
		case s := <-p.chatSubscription:
			mu.Lock()
			subscribers[s.ID] = s
			mu.Unlock() //fmt.Println("Subscribed to chat", subscribers)
		case m := <-p.chatAdded:
			for id, s := range subscribers {
				go func(id string, s *chatSubscription) {
					if m.RoomID != s.roomID {
						return // Skip this subscriber
					}
					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- m:
						fmt.Println("Sending message", id)
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}

	}
}

func (p *pubsub) Publish(message *dto.ChatMessage) {
	p.chatAdded <- message
}
