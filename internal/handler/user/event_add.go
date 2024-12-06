package user

import (
	"encoding/json"
	"healthy/internal/database/models"
	"healthy/internal/database/repositories"
)

func (u *UserHandler) AddUserEvents(req *AddUserEventReq) error {
	var userEvents = make([]*models.UserEvent, len(req.UserEvents))
	for idx, item := range req.UserEvents {
		bs, _ := json.Marshal(item.Content)
		userEvents[idx] = &models.UserEvent{
			Username:  req.Username,
			EventType: item.EventType,
			Content:   string(bs),
		}
	}

	err := repositories.AddUserEvents(u.db, userEvents)
	return err
}

type AddUserEventReq struct {
	UserEvents []*UserEvent `json:"user_events"`
	Username   string       `json:"username"`
}
