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

func TestSignIn(t *testing.T) {
	token, err := signIn(username, password)
	assert.Nil(t, err)
	assert.True(t, len(token) > 0)
}
func signIn(username, password string) (string, error) {
	type signInReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	req := &signInReq{
		Username: username,
		Password: password,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	resp, err := client.Post(host+"/api/v1/signin", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code: %d, msg: %s", resp.StatusCode, string(body))
	}
	type signInResp struct {
		Token string `json:"token"`
	}
	var respData signInResp
	if err := json.Unmarshal(body, &respData); err != nil {
		return "", fmt.Errorf("could not unmarshal response body: %w", err)
	}

	return respData.Token, nil
}
