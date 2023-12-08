package room

import (
	"github.com/GoodOnuii-B/seoltab-chat/config"
	"github.com/GoodOnuii-B/seoltab-chat/internal/dto"
	"github.com/GoodOnuii-B/seoltab-chat/tool/dynamo"
	"github.com/GoodOnuii-B/seoltab-chat/util"
)

type Service struct {
	dc *dynamo.Client
}

func NewService(cfg *config.AWSInfo) *Service {
	sess := dynamo.GetDynamoSession(cfg)
	dc := dynamo.GetDynamoClient(cfg, sess, cfg.Dynamo.Room.TableName)

	return &Service{
		dc: dc,
	}
}

func (s *Service) CreateRoom(
	groupId int32,
	lectureId int32,
	users []dto.User,
) (dto.Room, error) {
	var foundRoom *dto.Room
	for i := range users {
		var r *dto.RoomKey
		r = &dto.RoomKey{
			UserID:  users[i].Id,
			GroupID: groupId,
		}
		item, err := s.dc.GetItem(r)
		if err != nil {
			return dto.Room{}, err
		}

		if item != nil {
			var room dto.Room
			err := s.dc.ConvertToItem(item, &room)
			if err != nil {
				return dto.Room{}, err
			}

			foundRoom = &room
			continue
		}
		var tempID string
		if foundRoom == nil {
			UUID, err := util.CreateUUID()
			if err != nil {
				return dto.Room{}, err
			}
			ID := UUID.String()
			tempID = ID
		} else {
			tempID = foundRoom.ID
		}
		foundRoom = &dto.Room{
			ID:        tempID,
			GroupId:   groupId,
			LectureId: lectureId,
			Users:     users,
			UserId:    users[i].Id,
		}
		// !TODO: using dynamoDB batchWriteItem
		err = s.dc.PutItem(&foundRoom)
		if err != nil {
			return dto.Room{}, err
		}

	}

	return *foundRoom, nil
}

func (s *Service) GetRoom(
	userId string,
) ([]dto.Room, error) {
	//var rooms []dto.Room
	r, _, err := s.dc.ScanWithLimit("userId", userId, 100, nil)
	if err != nil {
		return nil, err
	}

	var rooms []dto.Room
	for i := range r {
		var room dto.Room
		err := s.dc.ConvertToItem(r[i], &room)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}
