# golang-bitbucket-cloud-build-status-notifier

[![Build Status](https://travis-ci.org/030/golang-bitbucket-cloud-build-status-notifier.svg?branch=master)](https://travis-ci.org/030/golang-bitbucket-cloud-build-status-notifier)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/golang-bitbucket-cloud-build-status-notifier)](https://goreportcard.com/report/github.com/030/golang-bitbucket-cloud-build-status-notifier)
[![dockeri.co](https://dockeri.co/image/utrecht/golang-bitbucket-cloud-build-status-notifier)](https://hub.docker.com/r/utrecht/golang-bitbucket-cloud-build-status-notifier)

Example for running this command given:

- **client id**: 1234
- **client secret**: 4567
- **build state**: INPROGRESS
- **git commit hash**: f76d123498a053c1789057a41d6c3fcvg8b49cd7
- **owner**: 030
- **repository name**: golang-bitbucket-cloud-build-status-notifier
- **build number**: 70
- **build url**: https://travis-ci.org/030/golang-bitbucket-cloud-build-status-notifier/builds/523263434

```
go run main.go -keyString 1234:5678 -buildState INPROGRESS -gitCommit f76d123498a053c1789057a41d6c3fcvg8b49cd7 -owner 030 -repositoryName golang-bitbucket-cloud-build-status-notifier -buildNumber 70 -buildURL https://travis-ci.org/030/golang-bitbucket-cloud-build-status-notifier/builds/523263434
```

To print the help message

```
docker run utrecht/golang-bitbucket-cloud-build-status-notifier:1.1.1 -h
```
