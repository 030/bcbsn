# golang-bitbucket-cloud-build-status-notifier

[![Build Status](https://travis-ci.org/030/golang-bitbucket-cloud-build-status-notifier.svg?branch=master)](https://travis-ci.org/030/golang-bitbucket-cloud-build-status-notifier)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/golang-bitbucket-cloud-build-status-notifier)](https://goreportcard.com/report/github.com/030/golang-bitbucket-cloud-build-status-notifier)
![Docker Pulls](https://img.shields.io/docker/pulls/utrecht/golang-bitbucket-cloud-build-status-notifier.svg)
![Issues](https://img.shields.io/github/issues-raw/030/golang-bitbucket-cloud-build-status-notifier.svg)
![Pull requests](https://img.shields.io/github/issues-pr-raw/030/golang-bitbucket-cloud-build-status-notifier.svg)
![Total downloads](https://img.shields.io/github/downloads/030/golang-bitbucket-cloud-build-status-notifier/total.svg)
![Latest Production Release Version](https://img.shields.io/github/release/030/golang-bitbucket-cloud-build-status-notifier.svg)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=030_golang-bitbucket-cloud-build-status-notifier&metric=bugs)](https://sonarcloud.io/dashboard?id=030_golang-bitbucket-cloud-build-status-notifier)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=030_golang-bitbucket-cloud-build-status-notifier&metric=code_smells)](https://sonarcloud.io/dashboard?id=030_golang-bitbucket-cloud-build-status-notifier)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=030_golang-bitbucket-cloud-build-status-notifier&metric=coverage)](https://sonarcloud.io/dashboard?id=030_golang-bitbucket-cloud-build-status-notifier)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=030_golang-bitbucket-cloud-build-status-notifier&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=030_golang-bitbucket-cloud-build-status-notifier)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=030_golang-bitbucket-cloud-build-status-notifier&metric=ncloc)](https://sonarcloud.io/dashboard?id=030_golang-bitbucket-cloud-build-status-notifier)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_golang-bitbucket-cloud-build-status-notifier&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=030_golang-bitbucket-cloud-build-status-notifier)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=030_golang-bitbucket-cloud-build-status-notifier&metric=alert_status)](https://sonarcloud.io/dashboard?id=030_golang-bitbucket-cloud-build-status-notifier)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_golang-bitbucket-cloud-build-status-notifier&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=030_golang-bitbucket-cloud-build-status-notifier)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=030_golang-bitbucket-cloud-build-status-notifier&metric=security_rating)](https://sonarcloud.io/dashboard?id=030_golang-bitbucket-cloud-build-status-notifier)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=030_golang-bitbucket-cloud-build-status-notifier&metric=sqale_index)](https://sonarcloud.io/dashboard?id=030_golang-bitbucket-cloud-build-status-notifier)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=030_golang-bitbucket-cloud-build-status-notifier&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=030_golang-bitbucket-cloud-build-status-notifier)

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

[![dockeri.co](https://dockeri.co/image/utrecht/golang-bitbucket-cloud-build-status-notifier)](https://hub.docker.com/r/utrecht/golang-bitbucket-cloud-build-status-notifier)

```
docker run utrecht/golang-bitbucket-cloud-build-status-notifier:1.1.1 -h
```
