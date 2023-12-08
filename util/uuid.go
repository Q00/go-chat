package util

import (
	"github.com/gofrs/uuid"
	"time"
)

func CreateUUID() (*uuid.UUID, error) {
	id, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func UUIDTime(ID *uuid.UUID) (time.Time, error) {
	ts, err := uuid.TimestampFromV6(*ID)
	if err != nil {
		return time.Time{}, err
	}

	t, err := ts.Time()
	if err != nil {
		return time.Time{}, err
	}
	return t, err
}
