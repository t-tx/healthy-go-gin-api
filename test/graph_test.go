package test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphGet(t *testing.T) {
	token, err := signIn(username, password)
	assert.Nil(t, err)

	err = graphGet(token)
	assert.Nil(t, err)
}
func graphGet(token string) error {
	httpReq, err := http.NewRequest("GET", host+"/api/v1/user/graph", nil)
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
	fmt.Println("graph", string(body))
	return nil

}