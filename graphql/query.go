package graphql

import (
	"context"
)

func (r *Resolver) Rooms(ctx context.Context, input struct{ UserId string }) ([]Room, error) {
	rooms, err := r.roomService.GetRoom(input.UserId)
	if err != nil {
		return nil, err
	}

	var gqlRooms []Room
	for _, room := range rooms {
		var gqlUsers []User
		for _, user := range room.Users {
			gqlUsers = append(gqlUsers, User{
				id:       user.Id,
				name:     user.Name,
				userType: user.UserType,
			})
		}
		gqlRoom := Room{
			id:        room.ID, // Ensure these fields are exported
			groupId:   room.GroupId,
			lectureId: room.LectureId,
			users:     gqlUsers,
		}

		gqlRooms = append(gqlRooms, gqlRoom)
	}

	return gqlRooms, nil
}

func (r *Resolver) ChatMessages(ctx context.Context, input struct {
	RoomId     string
	Limit      int32
	NextCursor *string
}) ChatMessagePage {
	chatMessages, lastCursor, err := r.chatService.GetChatMessages(input.RoomId, input.Limit, input.NextCursor)
	if err != nil {
		return ChatMessagePage{
			chatMessages: []ChatMessage{},
			nextCursor:   nil,
		}
	}

	var gqlChatMessages []ChatMessage
	for _, chatMessage := range chatMessages {
		gqlChatMessage := ChatMessage{
			id:     chatMessage.Id,
			roomId: chatMessage.RoomID,
			user: &User{
				id:       chatMessage.User.Id,
				name:     chatMessage.User.Name,
				userType: chatMessage.User.UserType,
			},
			chatMessageType: chatMessage.ChatMessageType,
			message:         chatMessage.Message,
			createdAt:       chatMessage.CreatedAt,
			isReport:        chatMessage.IsReport,
		}
		gqlChatMessages = append(gqlChatMessages, gqlChatMessage)
	}

	return ChatMessagePage{
		chatMessages: gqlChatMessages,
		nextCursor:   lastCursor,
	}
}
