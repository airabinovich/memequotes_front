package http

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	httpClient *http.Client
)

func init() {
	httpClient = &http.Client{}
}

// ResponseBodyAsString returns the a string representation of a response's body
func ResponseBodyAsString(resp *http.Response) (string, error) {
	if resp.Body == nil {
		return "", nil
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	return string(bytes), err
}

// RequestBodyAsString returns the a string representation of a requests's body
func RequestBodyAsString(req *http.Request) (string, error) {
	// TODO: validate this, body is never nil, no matter if request has nil body. Go by default always create an io.ReaderCloser
	if req.Body == nil {
		return "", nil
	}
	bytes, err := ioutil.ReadAll(req.Body)
	return string(bytes), err
}

// Retry function. It receives the function to be executed, how many times should be retried if there's an error
// and sleep time between retries
//  Retry( func() {
// 	 fmt.Print("Foo barr ")
//  }, 2, times.Duration(100 * time.Millisecond))
//
// Fn should be retried twice with a sleep of 100ms. time.
func Retry(fn func() error, times int, sleepDuration time.Duration) (err error) {
	for err = fn(); err != nil && times > 1; {
		time.Sleep(sleepDuration)

		times, err = times-1, fn()
	}
	return err
}

// ForwardHeaders takes headers from request context and add those headers to the given request.
func ForwardHeaders(ctx context.Context, req *http.Request) *http.Request {
	return req
}

// Client returns a new http client
func Client() *http.Client {
	return httpClient
}
