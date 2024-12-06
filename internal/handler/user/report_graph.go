package user

import (
	"errors"
	"healthy/internal/pkg/defined"
)

func (u *UserHandler) GetReportGraphs(req *GetReportGraphsReq) (map[string][]*Graph, error) {
	events, err := u.GetUserEvents(&GetUserEventsReq{Username: req.Username})
	if err != nil {
		return nil, errors.New("could not get user events")
	}

	var result = make(map[string][]*Graph, 2)
	for _, event := range events {
		if event.EventType != defined.MEASURE {
			continue
		}
		switch event.Content["type"] {
		case defined.WEIGHT:
			result[defined.WEIGHT] = append(result[defined.WEIGHT], &Graph{
				Time:  event.CreatedAt,
				Value: event.Content["value"].(float64),
			})
		case defined.HEIGHT:
			result[defined.HEIGHT] = append(result[defined.WEIGHT], &Graph{
				Time:  event.CreatedAt,
				Value: event.Content["value"].(float64),
			})
		}
	}

	return result, nil
}

type Graph struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}
type GetReportGraphsReq struct {
	Username string `json:"username"`
}
