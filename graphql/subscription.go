package graphql

import (
	"context"
	"fmt"
)

func (r *Resolver) ChatMessage(ctx context.Context, input struct{ RoomId string }) <-chan *ChatMessage {
	ID, c := r.pubsub.Subscribe(input.RoomId)
	chatMessageChan := make(chan *ChatMessage)

	go func() {
		for chatAdded := range c {
			// Assuming dto.ChatMessage can be converted to ChatMessage
			chatMessage := &ChatMessage{
				id:     chatAdded.Id,
				roomId: chatAdded.RoomID,
				user: &User{
					id:       chatAdded.User.Id,
					name:     chatAdded.User.Name,
					userType: chatAdded.User.UserType,
				},
				chatMessageType: chatAdded.ChatMessageType,
				message:         chatAdded.Message,
				createdAt:       chatAdded.CreatedAt,
				isReport:        chatAdded.IsReport,
			}

			chatMessageChan <- chatMessage
		}

	}()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ID, "Stopping")
			r.pubsub.UnSubscribe(ID)
		}
	}()

	return chatMessageChan
}
