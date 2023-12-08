package dto

import "time"

type ChatMessage struct {
	Id              string    `json:"id"`
	RoomID          string    `json:"roomId"`
	User            User      `json:"user"`
	Message         string    `json:"message"`
	ChatMessageType string    `json:"chatMessageType"` // MESSAGE, EMOJI, or IMAGE
	IsReport        *bool     `json:"isReport"`
	CreatedAt       time.Time `json:"createdAt"`
}
