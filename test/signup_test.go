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

func TestSignUp(t *testing.T) {
	err := signUp(username, password, "male")
	assert.Nil(t, err)
}

func signUp(username, password, gender string) error {
	type signUpReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Gender   string `json:"gender"`
	}
	req := &signUpReq{
		Username: username,
		Password: password,
		Gender:   gender,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.Post(host+"/api/v1/signup", "application/json", bytes.NewBuffer(reqBody))
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
