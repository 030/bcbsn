package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/030/go-utils/utils"
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

	buildStatus(oauthdance.Bearer(*keyString), *buildState, *gitCommit, *owner, *repositoryName, *buildNumber, *buildURL)
}
