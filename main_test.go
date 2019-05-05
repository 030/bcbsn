package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHTTPClientShouldTimeOutAfterTenSeconds(t *testing.T) {
	expectedError := "net/http: request canceled (Client.Timeout exceeded while awaiting headers)"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 11)
	}))
	defer svr.Close()

	err := buildStatusImpl(svr.URL, "bearer", "state", "c1o2m3m4i5t6", "owner", "repoOwner", "1", "http://build.url")

	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Error actual = %v and expected = %v", err, expectedError)
	}
}
