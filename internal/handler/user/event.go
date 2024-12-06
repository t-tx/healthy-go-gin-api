package user

type UserEvent struct {
	EventType string                 `json:"event_type"`
	Content   map[string]interface{} `json:"content"`
}
