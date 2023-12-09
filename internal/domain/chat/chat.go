package chat

import (
	"github.com/Q00/go-chat/config"
	"github.com/Q00/go-chat/internal/dto"
	"github.com/Q00/go-chat/tool/dynamo"
	"github.com/Q00/go-chat/util"
	"github.com/gofrs/uuid"
)

type Service struct {
	dc *dynamo.Client
}

func NewService(cfg *config.AWSInfo) *Service {
	sess := dynamo.GetDynamoSession(cfg)
	dc := dynamo.GetDynamoClient(cfg, sess, cfg.Dynamo.Chat.TableName)

	return &Service{
		dc: dc,
	}
}

func (s *Service) AddChatMessage(
	message dto.ChatMessage,
	pubsub PUBSUB,
) (bool, error) {
	ID, err := util.CreateUUID()
	if err != nil {
		return false, err
	}
	message.Id = ID.String()
	t, err := util.UUIDTime(ID)
	if err != nil {
		return false, err
	}
	message.CreatedAt = t.UTC()

	err = s.dc.PutItem(message)
	if err != nil {
		return false, err
	}

	pubsub.Publish(&message)

	return true, nil
}

// GetChatMessages returns a list of chat messages for a given room
func (s *Service) GetChatMessages(
	roomId string,
	limit int32,
	nextCursor *string,
) ([]dto.ChatMessage, *string, error) {

	var p *dto.ChatMessageKey
	if nextCursor != nil {
		p = &dto.ChatMessageKey{
			ID:     *nextCursor,
			RoomID: roomId,
		}
	}

	c, lastCursor, err := s.dc.ScanWithLimit("roomId", roomId, limit, p)
	if err != nil {
		return nil, nil, err
	}
	var messages []dto.ChatMessage
	for _, item := range c {
		message := dto.ChatMessage{}
		err = s.dc.ConvertToItem(item, &message)
		if err != nil {
			return nil, nil, err
		}
		id, err := uuid.FromString(message.Id)
		if err != nil {
			return nil, nil, err
		}
		createdAt, err := util.UUIDTime(&id)
		message.CreatedAt = createdAt.UTC()
		messages = append(messages, message)
	}

	return messages, lastCursor, nil
}

func (s *Service) ReportChatMessage(chatMessageID string, roomID string, IsReport bool) (bool, error) {
	var p *dto.ChatMessageKey
	p = &dto.ChatMessageKey{
		ID:     chatMessageID,
		RoomID: roomID,
	}
	item, err := s.dc.GetItem(p)
	if err != nil {
		return false, err
	}

	var cm dto.ChatMessage
	err = s.dc.ConvertToItem(item, &cm)
	if err != nil {
		return false, err
	}

	cm.IsReport = &IsReport

	err = s.dc.PutItem(cm)
	if err != nil {
		return false, err
	}

	return true, nil
}
