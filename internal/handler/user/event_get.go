package user

import (
	"encoding/json"
	"healthy/internal/database/repositories"
	"healthy/internal/pkg/utils"
)

func (u *UserHandler) GetUserEvents(req *GetUserEventsReq) ([]*GetUserEventsResp, error) {
	userEvents, err := repositories.ListUserEvent(u.db, req.Username)
	if err != nil {
		return nil, err
	}
	var output = make([]*GetUserEventsResp, len(userEvents))
	for idx, event := range userEvents {
		var content = make(map[string]interface{})
		err := json.Unmarshal([]byte(event.Content), &content)
		if err != nil {
			return nil, err
		}
		output[idx] = &GetUserEventsResp{
			EventType: event.EventType,
			Content:   content,
			CreatedAt: utils.GetTimeString(event.CreatedAt),
		}
	}
	return output, nil
}

type GetUserEventsReq struct {
	Username string `json:"username"`
}
type GetUserEventsResp struct {
	EventType string                 `json:"event_type"`
	Content   map[string]interface{} `json:"content"`
	CreatedAt string                 `json:"created_at"`
}
