# bcbsn

[![Build Status](https://travis-ci.org/030/bcbsn.svg?branch=master)](https://travis-ci.org/030/bcbsn)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/bcbsn)](https://goreportcard.com/report/github.com/030/bcbsn)
![DevOps SE Questions](https://img.shields.io/stackexchange/devops/t/bcbsn.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/utrecht/bcbsn.svg)
![Issues](https://img.shields.io/github/issues-raw/030/bcbsn.svg)
![Pull requests](https://img.shields.io/github/issues-pr-raw/030/bcbsn.svg)
![Total downloads](https://img.shields.io/github/downloads/030/bcbsn/total.svg)
![License](https://img.shields.io/github/license/030/bcbsn.svg)
![Repository Size](https://img.shields.io/github/repo-size/030/bcbsn.svg)
![Contributors](https://img.shields.io/github/contributors/030/bcbsn.svg)
![Commit activity](https://img.shields.io/github/commit-activity/m/030/bcbsn.svg)
![Last commit](https://img.shields.io/github/last-commit/030/bcbsn.svg)
![Release date](https://img.shields.io/github/release-date/030/bcbsn.svg)
![Latest Production Release Version](https://img.shields.io/github/release/030/bcbsn.svg)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=030_bcbsn&metric=bugs)](https://sonarcloud.io/dashboard?id=030_bcbsn)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=030_bcbsn&metric=code_smells)](https://sonarcloud.io/dashboard?id=030_bcbsn)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=030_bcbsn&metric=coverage)](https://sonarcloud.io/dashboard?id=030_bcbsn)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=030_bcbsn&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=030_bcbsn)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=030_bcbsn&metric=ncloc)](https://sonarcloud.io/dashboard?id=030_bcbsn)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_bcbsn&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=030_bcbsn)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=030_bcbsn&metric=alert_status)](https://sonarcloud.io/dashboard?id=030_bcbsn)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_bcbsn&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=030_bcbsn)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=030_bcbsn&metric=security_rating)](https://sonarcloud.io/dashboard?id=030_bcbsn)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=030_bcbsn&metric=sqale_index)](https://sonarcloud.io/dashboard?id=030_bcbsn)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=030_bcbsn&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=030_bcbsn)

The Bitbucket Cloud Build Status Notifier (BCBSN) is a tool that sends a build status to bitbucket cloud.
Oauth credentials are required to perform this action and these have to be created in the user settings
menu that resides in bitbucket cloud. Note: a Callback URL, e.g. https://ci-system-name is required in order
to create Oauth credentials.

Example for running this command given:

- **client id**: 1234
- **client secret**: 4567
- **build state**: INPROGRESS
- **git commit hash**: f76d123498a053c1789057a41d6c3fcvg8b49cd7
- **owner**: 030
- **repository name**: bcbsn
- **build number**: 70
- **build url**: https://travis-ci.org/030/bcbsn/builds/523263434

```
./bcbsn -keyString 1234:5678 -buildState INPROGRESS -gitCommit f76d123498a053c1789057a41d6c3fcvg8b49cd7 -owner 030 -repositoryName bcbsn -buildNumber 70 -buildURL https://travis-ci.org/030/bcbsn/builds/523263434
```

[![dockeri.co](https://dockeri.co/image/utrecht/bcbsn)](https://hub.docker.com/r/utrecht/bcbsn)

To print the help message

```
docker run utrecht/bcbsn:2.0.0 -h
```
