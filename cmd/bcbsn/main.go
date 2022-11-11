package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const bitbucketAPIEndpoint = "https://api.bitbucket.org/2.0/repositories"

type Body struct {
	UUID  string `json:"uuid"`  // The Git commit hash
	Key   string `json:"key"`   // The Build ID
	URL   string `json:"url"`   // The URL of the specific build
	State string `json:"state"` // The state of the build
	Name  string `json:"name"`  // The identifier of the build
}

var httpClientWithTimeout = http.Client{
	Timeout: time.Second * 10,
}

func prepareBody(commit, key, url, state, name string) ([]byte, error) {
	b := Body{commit, key, url, state, name}
	return json.Marshal(b)
}

func prepareOAUTH2HTTPClient(httpClient http.Client, clientID, clientSecret string) *http.Client {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://bitbucket.org/site/oauth2/access_token",
	}

	// Instead of the default httpClient use our custom client (which sets a timeout)
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, httpClient)
	return config.Client(ctx)
}

// We need to be able to overwrite the endpoint in our unit tests
// So we use the approach as described here: https://stackoverflow.com/a/33774754/10224882
func setBuildStatus(clientID, clientSecret, owner, repoSlug, commit, key, url, state, name string) error {
	httpClient := prepareOAUTH2HTTPClient(httpClientWithTimeout, clientID, clientSecret)
	return setBuildStatusImpl(*httpClient, bitbucketAPIEndpoint, owner, repoSlug, commit, key, url, state, name)
}

func setBuildStatusImpl(httpClient http.Client, bitbucketEndpoint, owner, repoSlug, commit, key, url, state, name string) error {
	body, err := prepareBody(commit, key, url, state, name)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", bitbucketEndpoint+"/"+owner+"/"+repoSlug+"/commit/"+commit+"/statuses/build", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if resp != nil {
		// Bitbucket returns a 201 when it first creates the build badge all subsequent
		// requests (i.e. updates on the same build) return a 200. We can't write
		// it as r != b || r != c as go vet rejects this, see https://github.com/golang/go/issues/28446
		if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
			return fmt.Errorf("Expected 200 or 201, but got %s", resp.Status)
		}
		defer resp.Body.Close()
	}

	if err != nil {
		return err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))

	return nil
}

func main() {
	var clientID = flag.String(
		"clientID",
		"",
		"The clientID used for the 'client credentials' token flow with BitBucket")

	var clientSecret = flag.String(
		"clientSecret",
		"",
		"The clientSecret used for the 'client credentials' token flow with BitBucket")

	var state = flag.String(
		"state",
		"",
		"The state, e.g. SUCCESSFUL, INPROGRESS or FAILED")

	var commit = flag.String(
		"commit",
		"",
		"The commit, e.g. 57484fd5460017aef111e8b4ec116a30ff0b4904")

	var owner = flag.String(
		"owner",
		"",
		"The owner of the repository, e.g. it is 'atlassian' in 'https://bitbucket.org/atlassian/stash-example-plugin/src/master/'")

	var repoSlug = flag.String(
		"repoSlug",
		"",
		"The repoSlug, e.g. some-repository")

	var key = flag.String(
		"key",
		"",
		"The key, e.g. a unique id of the build (use the build id)")

	var url = flag.String(
		"url",
		"",
		"The url, e.g. https://travis-ci.org/030/bcbsn/builds/523263434")

	var name = flag.String(
		"name",
		"",
		"An identifier for the build e.g. 'build 2'")

	flag.Parse()

	if *clientID == "" || *clientSecret == "" || *state == "" || *commit == "" || *owner == "" || *repoSlug == "" || *key == "" || *url == "" || *name == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := setBuildStatus(*clientID, *clientSecret, *owner, *repoSlug, *commit, *key, *url, *state, *name); err != nil {
		log.Fatal(err)
	}
}
