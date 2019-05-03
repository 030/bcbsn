package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/030/golang-bitbucket-cloud-oauth-dance/oauthdance"
	log "github.com/sirupsen/logrus"
)

func buildStatus(bearer string, buildState string, gitCommit string, owner string, repositoryOwner string, buildNumber string, buildURL string) {
	body := strings.NewReader(`{"key":"` + buildNumber + `","state":"` + buildState + `","url":"` + buildURL + `"}`)
	req, err := http.NewRequest("POST", "https://api.bitbucket.org/2.0/repositories/"+owner+"/"+repositoryOwner+"/commit/"+gitCommit+"/statuses/build", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+bearer)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
}

func main() {
	var keyString = flag.String(
		"keyString",
		"",
		"The concatenated clientId (key in bitbucket terms) and clientSecret (secret in bitbucket terms), e.g. clientId:clientSecret")

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
		"The owner of the repository, e.g. it is 'atlassian' in 'https://bitbucket.org/atlassian/stash-example-plugin/src/master/'")

	var repositoryName = flag.String(
		"repositoryName",
		"",
		"The name of the repository, e.g. it is 'stash-example-plugin' in 'https://bitbucket.org/atlassian/stash-example-plugin/src/master/'")

	var buildNumber = flag.String(
		"buildNumber",
		"",
		"The build number, e.g. the unique identifier of a build e.g. 333")

	var buildURL = flag.String(
		"buildURL",
		"",
		"The url of the specific build (so bitbucket can direct you back), e.g. https://travis-ci.org/030/golang-bitbucket-cloud-build-status-notifier/jobs/523263435")

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "golang-bitbucket-cloud-build-status-notifier sets the status of a build in bitbucket.\n\n")
		flag.PrintDefaults()
	}

	if *keyString == "" || *buildState == "" || *gitCommit == "" || *owner == "" || *repositoryName == "" || *buildNumber == "" || *buildURL == "" {
		flag.Usage()
		os.Exit(1)
	}

	accessToken, err := oauthdance.Bearer(*keyString)
	if err != nil {
		log.Fatal(err)
	}

	buildStatus(accessToken, *buildState, *gitCommit, *owner, *repositoryName, *buildNumber, *buildURL)
}
