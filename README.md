# golang-bitbucket-cloud-build-status-notifier

[![Build Status](https://travis-ci.org/030/golang-bitbucket-cloud-build-status-notifier.svg?branch=master)](https://travis-ci.org/030/golang-bitbucket-cloud-build-status-notifier)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/golang-bitbucket-cloud-build-status-notifier)](https://goreportcard.com/report/github.com/030/golang-bitbucket-cloud-build-status-notifier)

```
go run main.go -keyString key:string -buildState INPROGRESS -gitCommit f76d123498a053c1789057a41d6c3fcvg8b49cd7 -owner owner -repositoryName repoName -buildNumber 13 -buildURL jobUrl
```

or

```
docker run utrecht/golang-bitbucket-cloud-build-status-notifier:1.0.0 -h
```
