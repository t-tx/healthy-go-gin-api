package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"healthy/internal/pkg/defined"
	"io"
	"math/rand"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type event struct {
	EventType string                 `json:"event_type"`
	Content   map[string]interface{} `json:"content"`
}

func newEvent(eventType string, kv ...any) *event {
	content := make(map[string]interface{})
	for i := 0; i < len(kv); i += 2 {
		content[kv[i].(string)] = kv[i+1]
	}
	return &event{
		EventType: eventType,
		Content:   content,
	}
}

func generateRandomEvents() []*event {
	eventTypes := []string{defined.EXERCISE, "meal", defined.MEASURE}
	exercises := []string{"running", "cycling", "swimming"}
	meals := [][]map[string]interface{}{
		{
			{"food": "egg", "quantity": 2},
			{"food": "sashimi", "quantity": 10},
		},
		{
			{"food": "egg", "quantity": 2},
			{"food": "chicken", "quantity": 1},
		},
		{
			{"food": "pho", "quantity": 1},
		},
	}
	measurements := []string{defined.HEIGHT, defined.WEIGHT}

	numEvents := rand.Intn(11) + 20
	events := make([]*event, numEvents)

	for i := 0; i < numEvents; i++ {
		eventType := eventTypes[rand.Intn(len(eventTypes))]
		switch eventType {
		case defined.EXERCISE:
			events[i] = newEvent(defined.EXERCISE, "type", exercises[rand.Intn(len(exercises))], "duration", rand.Intn(90)+10)

		case defined.MEASURE:
			measureType := measurements[rand.Intn(len(measurements))]
			events[i] = newEvent(defined.MEASURE, "type", measureType, "value", rand.Intn(50)+50)

		case "meal":
			mealType := []string{"breakfast", "lunch", "dinner"}[rand.Intn(3)]
			events[i] = newEvent("meal", "type", mealType, "dish", meals[rand.Intn(len(meals))])
		default:
			panic("unknown event type")
		}
	}

	return events
}
func TestEventAdd(t *testing.T) {
	token, err := signIn(username, password)
	assert.Nil(t, err)

	var req = generateRandomEvents()
	err = eventAdd(token, req)

	assert.Nil(t, err)
}
func eventAdd(token string, req []*event) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest("POST", host+"/api/v1/user/events", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:  "token",
		Value: token,
		Path:  "/",
	}
	httpReq.AddCookie(cookie)

	resp, err := client.Do(httpReq)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d, msg: %s", resp.StatusCode, string(body))
	}

	return nil

}
