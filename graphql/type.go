package graphql

import (
	graphqlGo "github.com/graph-gophers/graphql-go"
	"time"
)

type Room struct {
	id        string `json:"id"`
	groupId   int32  `json:"groupId"`
	lectureId int32  `json:"lectureId"`
	users     []User `json:"users"`
}

func (r Room) Id() graphqlGo.ID {
	return graphqlGo.ID(r.id)
}

func (r Room) GroupId() int32 {
	return r.groupId
}

func (r Room) LectureId() int32 {
	return r.lectureId
}

func (r Room) Users() []User {
	return r.users
}

type UserInput struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	UserType string `json:"userType"` // STUDENT or TEACHER
}

type User struct {
	id       string `json:"id"`
	name     string `json:"name"`
	userType string `json:"userType"` // STUDENT or TEACHER
}

func (u User) Id() graphqlGo.ID {
	return graphqlGo.ID(u.id)
}

func (u User) Name() string {
	return u.name
}

func (u User) UserType() string {
	return u.userType
}

type ChatMessageInput struct {
	RoomId          string     `json:"roomId"`
	User            *UserInput `json:"user"`
	Message         string     `json:"message"`
	ChatMessageType string     `json:"chatMessageType"` // MESSAGE, EMOJI, or IMAGE
}

type ChatMessagePage struct {
	chatMessages []ChatMessage `json:"chatMessages"`
	nextCursor   *string       `json:"nextCursor"`
}

func (c ChatMessagePage) ChatMessages() []ChatMessage {
	return c.chatMessages
}

func (c ChatMessagePage) NextCursor() *string {
	return c.nextCursor
}

type ChatMessage struct {
	id              string    `json:"id"`
	roomId          string    `json:"roomId"`
	user            *User     `json:"user"`
	message         string    `json:"message"`
	chatMessageType string    `json:"chatMessageType"` // MESSAGE, EMOJI, or IMAGE
	isReport        *bool     `json:"isReport"`
	createdAt       time.Time `json:"createdAt"`
}

func (c ChatMessage) Id() graphqlGo.ID {
	return graphqlGo.ID(c.id)
}

func (c ChatMessage) RoomId() string {
	return c.roomId
}

func (c ChatMessage) User() *User {
	return c.user
}

func (c ChatMessage) Message() string {
	return c.message
}

func (c ChatMessage) ChatMessageType() string {
	return c.chatMessageType
}

func (c ChatMessage) IsReport() *bool {
	return c.isReport
}

func (c ChatMessage) CreatedAt() graphqlGo.Time {
	return graphqlGo.Time{Time: c.createdAt}
}
