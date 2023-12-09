package graphql

import (
	"github.com/Q00/go-chat/internal/dto"
)

func (r *Resolver) AddChatMessage(input struct {
	Message ChatMessageInput
}) (bool, error) {
	user := dto.User{
		Id:       input.Message.User.Id,
		Name:     input.Message.User.Name,
		UserType: input.Message.User.UserType,
	}
	message := dto.ChatMessage{
		RoomID:          input.Message.RoomId,
		User:            user,
		ChatMessageType: input.Message.ChatMessageType,
		Message:         input.Message.Message,
	}

	return r.chatService.AddChatMessage(message, r.pubsub)
}

func (r *Resolver) CreateRoom(input struct {
	GroupId   int32
	LectureId int32
	Users     []UserInput
}) (Room, error) {
	var users []dto.User
	for _, userInput := range input.Users {
		users = append(users, dto.User{
			Id:       userInput.Id,
			Name:     userInput.Name,
			UserType: userInput.UserType,
		})
	}

	room, err := r.roomService.CreateRoom(input.GroupId, input.LectureId, users)
	if err != nil {
		return Room{}, err
	}

	// Convert back to GraphQL types if needed
	var gqlUsers []User
	for _, user := range room.Users {
		gqlUsers = append(gqlUsers, User{
			id:       user.Id,
			name:     user.Name,
			userType: user.UserType,
		})
	}

	// Construct the GraphQL Room
	gqlRoom := Room{
		id:        room.ID, // Ensure these fields are exported
		lectureId: room.LectureId,
		groupId:   room.GroupId,
		users:     gqlUsers,
	}

	return gqlRoom, nil
}

func (r *Resolver) ReportChatMessage(input struct {
	ChatMessageID string
	RoomID        string
	IsReport      bool
}) (bool, error) {
	r.chatService.ReportChatMessage(input.ChatMessageID, input.RoomID, input.IsReport)
	return true, nil
}
