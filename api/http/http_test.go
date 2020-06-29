package http

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequestBodyAsString(t *testing.T) {
	t.Log("Body should be parsed as string successfully")
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"key":"value","int-key":123}`))
	s, err := RequestBodyAsString(req)

	assert.NoError(t, err)
	assert.Equal(t, `{"key":"value","int-key":123}`, s)
}

func TestWithoutRequestBody(t *testing.T) {
	t.Log("Empty request body should be parsed successfully")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	s, err := RequestBodyAsString(req)

	assert.NoError(t, err)
	assert.Equal(t, "", s)
}

func TestResponseBodyAsString(t *testing.T) {
	t.Log("Response body should be parsed successfully")
	respRecorder := httptest.NewRecorder()
	_, _ = respRecorder.WriteString(`{"key":"value","int-key":7384}`)

	s, err := ResponseBodyAsString(respRecorder.Result())

	assert.NoError(t, err)
	assert.Equal(t, `{"key":"value","int-key":7384}`, s)
}

func TestWithoutResponse(t *testing.T) {
	t.Log("Empty response body should be parsed successfully")
	respRecorder := httptest.NewRecorder()

	s, err := ResponseBodyAsString(respRecorder.Result())

	assert.Nil(t, err)
	assert.Equal(t, "", s)
}

func TestRetryFunction(t *testing.T) {
	t.Log("retry function should be executed 3 times with a sleep of 100 millis. Should return an error")

	count := 0
	err := Retry(func() error {
		count++
		return errors.New("some error")
	}, 3, time.Duration(1000000))

	assert.Error(t, err)
	assert.Equal(t, 3, count)
}

func TestRetryFunctionSuccessCall(t *testing.T) {
	t.Log("retry function should be executed 2 times with a sleep of 100 millis. Should return no error")

	count := 0
	err := Retry(func() error {
		if count < 2 {
			count++
			return errors.New("some error")
		}
		return nil

	}, 3, time.Duration(1000000))

	assert.NoError(t, err)
	assert.Equal(t, 2, count)
}

func TestForwardHeaders(t *testing.T) {
	t.Log("Custom headers should be forwarded")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("custom-header", "some value")
	req.Header.Set("other-header", "other value")

	assert.Equal(t, "some value", req.Header.Get("custom-header"))
	assert.Equal(t, "other value", req.Header.Get("other-header"))
}

func TestHttpClient(t *testing.T) {
	t.Log("Client should return an httpClient")

	client := Client()

	assert.NotNil(t, client)
}
