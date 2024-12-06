package test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArticlesGet(t *testing.T) {
	err := articleGet()
	assert.Nil(t, err)
}
func articleGet() error {
	httpReq, err := http.NewRequest("GET", host+"/api/v1/articles/health", nil)
	if err != nil {
		return err
	}

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
	fmt.Println("article", string(body))
	return nil

}
