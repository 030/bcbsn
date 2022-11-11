package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

const (
	commit   = "c1o2m3m4i5t6"
	key      = "1"
	url      = "http://build.url"
	state    = "SUCCESSFUL"
	name     = "build #1"
	owner    = "030"
	repoSlug = "repo_slug"
)

var testHTTPClient = http.Client{
	Timeout: time.Second,
}

func TestHTTPClientShouldTimeOutAfterTenSeconds(t *testing.T) {
	expectedError := "context deadline exceeded (Client.Timeout exceeded while awaiting headers)"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 2)
	}))
	defer svr.Close()

	err := setBuildStatusImpl(testHTTPClient, svr.URL, owner, repoSlug, commit, key, url, state, name)

	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Expected '%v', but got '%v'", expectedError, err)
	}
}

func TestHTTPClientShouldFailIfResponseIsNotA201(t *testing.T) {
	expectedError := "Expected 200 or 201, but got 400 Bad Request"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer svr.Close()

	err := setBuildStatusImpl(testHTTPClient, svr.URL, owner, repoSlug, commit, key, url, state, name)

	if err.Error() != expectedError {
		t.Errorf("Expected '%v', but got '%v'", expectedError, err)
	}
}

func TestHTTPClientShouldAcceptAnHTTP200(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer svr.Close()

	err := setBuildStatusImpl(testHTTPClient, svr.URL, owner, repoSlug, commit, key, url, state, name)
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
}

func TestHTTPClientShouldAcceptAnHTTP201(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer svr.Close()

	err := setBuildStatusImpl(testHTTPClient, svr.URL, owner, repoSlug, commit, key, url, state, name)
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
}

func TestPostBodyHasExpectedContent(t *testing.T) {
	expectedURL := "/030/repo_slug/commit/c1o2m3m4i5t6/statuses/build"
	expectedBody := Body{commit, key, url, state, name}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != "POST" {
			t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
		}
		if r.URL.EscapedPath() != expectedURL {
			t.Errorf("Expected request to '%s', got '%s'", expectedURL, r.URL.EscapedPath())
		}
		if r.Header.Get("Content-type") != "application/json" {
			t.Errorf("Expected content-type to be 'application/json', but got '%s'", r.Header.Get("Content-type"))
		}

		decoder := json.NewDecoder(r.Body)
		var actualBody Body
		err := decoder.Decode(&actualBody)
		if err != nil {
			t.Errorf("Error while reading request JSON: %s", err)
		}

		if !reflect.DeepEqual(expectedBody, actualBody) {
			t.Errorf("Expected request body '%s', but was '%s'", expectedBody, actualBody)
		}
	}))
	defer svr.Close()

	if err := setBuildStatusImpl(testHTTPClient, svr.URL, owner, repoSlug, commit, key, url, state, name); err != nil {
		log.Fatal(err)
	}
}
