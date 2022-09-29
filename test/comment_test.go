// +e2e

package test

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BaseUrl + "/api/comment")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/", "author": "charles", "body": "hello world"}`).
		Post(BaseUrl + "/api/comment")
	if err != nil {
		t.Fail()
	}

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
