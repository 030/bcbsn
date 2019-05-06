package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Body struct {
	UUID  string `json:"uuid"`  // The Git commit hash
	Key   string `json:"key"`   // The Build ID
	URL   string `json:"url"`   // The URL of the specific build
	State string `json:"state"` // The state of the build
	Name  string `json:"name"`  // The identifier of the build
}

var httpClient = http.Client{
	Timeout: time.Second * 10,
}

const bitbucketAPIEndpoint = "https://api.bitbucket.org/2.0/repositories"

func prepareBody(commit, key, url, state, name string) ([]byte, error) {
	b := Body{commit, key, url, state, name}
	return json.Marshal(b)
}

// We need to be able to overwrite the endpoint in our unit tests
// So we use the approach as described here: https://stackoverflow.com/a/33774754/10224882
func buildStatus(owner, repoSlug, commit, key, url, state, name string) error {
	return buildStatusImpl(httpClient, bitbucketAPIEndpoint, owner, repoSlug, commit, key, url, state, name)
}

func buildStatusImpl(httpClient http.Client, bitbucketEndpoint, owner, repoSlug, commit, key, url, state, name string) error {
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
		if resp.StatusCode != http.StatusCreated {
			return fmt.Errorf("Expected %d, but got %s", http.StatusCreated, resp.Status)
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
	var state = flag.String(
		"buildState",
		"",
		"The buildState, e.g. SUCCESSFUL, INPROGRESS or FAILED")

	var commit = flag.String(
		"gitCommit",
		"",
		"The gitCommit, e.g. 57484fd5460017aef111e8b4ec116a30ff0b4904")

	var owner = flag.String(
		"owner",
		"",
		"The owner, e.g. your-name")

	var repoSlug = flag.String(
		"repositoryName",
		"",
		"The repositoryName, e.g. some-repository")

	var key = flag.String(
		"buildNumber",
		"",
		"The buildNumber, e.g. buildNumber")

	var url = flag.String(
		"buildURL",
		"",
		"The buildURL, e.g. buildURL")

	var name = flag.String(
		"name",
		"",
		"An identifier for the build")

	err := buildStatus(*owner, *repoSlug, *commit, *key, *url, *state, *name)
	if err != nil {
		log.Fatal(err)
	}
}
