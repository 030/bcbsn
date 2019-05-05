package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/030/go-utils/utils"
	"github.com/030/golang-bitbucket-cloud-oauth-dance/oauthdance"
	log "github.com/sirupsen/logrus"
)

type Body struct {
	State       string
	Key         string
	Name        string
	URL         string
	Description string
}

var netClient = http.Client{
	Timeout: time.Second * 10,
}

const bitbucketAPIEndpoint = "https://api.bitbucket.org/2.0/repositories"

// We need to be able to overwrite the endpoint in our unit tests
// So we use the approach as described here: https://stackoverflow.com/a/33774754/10224882
func buildStatus(bearer, buildState, gitCommit, owner, repositoryOwner, buildNumber, buildURL string) error {
	return buildStatusImpl(bitbucketAPIEndpoint, bearer, buildState, gitCommit, owner, repositoryOwner, buildNumber, buildURL)
}

func buildStatusImpl(bitbucketEndpoint, bearer, buildState, gitCommit, owner, repositoryOwner, buildNumber, buildURL string) error {
	body := strings.NewReader(`{"key":"` + buildNumber + `","state":"` + buildState + `","url":"` + buildURL + `"}`)
	req, err := http.NewRequest("POST", bitbucketEndpoint+"/"+owner+"/"+repositoryOwner+"/commit/"+gitCommit+"/statuses/build", body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+bearer)
	req.Header.Set("Content-Type", "application/json")

	resp, err := netClient.Do(req)
	if resp != nil {
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
	var keyString = flag.String(
		"keyString",
		"",
		"The concatenate key and string, e.g. key:string")

	var buildState = flag.String(
		"buildState",
		"",
		"The buildState, e.g. SUCCESSFUL, INPROGRESS or FAILED")

	var gitCommit = flag.String(
		"gitCommit",
		"",
		"The gitCommit, e.g. 57484fd5460017aef111e8b4ec116a30ff0b4904")

	var owner = flag.String(
		"owner",
		"",
		"The owner, e.g. your-name")

	var repositoryName = flag.String(
		"repositoryName",
		"",
		"The repositoryName, e.g. some-repository")

	var buildNumber = flag.String(
		"buildNumber",
		"",
		"The buildNumber, e.g. buildNumber")

	var buildURL = flag.String(
		"buildURL",
		"",
		"The buildURL, e.g. buildURL")

	utils.Debug()

	accessToken, err := oauthdance.Bearer(*keyString)
	if err != nil {
		log.Fatal(err)
	}

	buildStatus(accessToken, *buildState, *gitCommit, *owner, *repositoryName, *buildNumber, *buildURL)
}
