package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type diaryAddReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func TestDiaryAdd(t *testing.T) {
	token, err := signIn(username, password)
	assert.Nil(t, err)

	var req = []*diaryAddReq{
		{Title: "Test",
			Content: "Test",
		},
	}
	err = diaryAdd(token, req)
	assert.Nil(t, err)
}
func diaryAdd(token string, req []*diaryAddReq) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}
	httpReq, err := http.NewRequest("POST", host+"/api/v1/user/diaries", bytes.NewBuffer(reqBody))
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
