# bcbsn

[![CI](https://github.com/030/bcbsn/workflows/Go/badge.svg?event=push)](https://github.com/030/bcbsn/actions?query=workflow%3AGo)
[![GoDoc Widget]][godoc]
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/030/bcbsn?logo=go)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/bcbsn)](https://goreportcard.com/report/github.com/030/bcbsn)
[![StackOverflow SE Questions](https://img.shields.io/stackexchange/stackoverflow/t/bcbsn.svg?logo=stackoverflow)](https://stackoverflow.com/tags/bcbsn)
[![DevOps SE Questions](https://img.shields.io/stackexchange/devops/t/bcbsn.svg?logo=stackexchange)](https://devops.stackexchange.com/tags/bcbsn)
[![ServerFault SE Questions](https://img.shields.io/stackexchange/serverfault/t/bcbsn.svg?logo=serverfault)](https://serverfault.com/tags/bcbsn)
[![Docker Pulls](https://img.shields.io/docker/pulls/utrecht/bcbsn?logo=docker&logoColor=white)](https://hub.docker.com/r/utrecht/bcbsn)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/utrecht/bcbsn?logo=docker&logoColor=white&sort=semver)
![Issues](https://img.shields.io/github/issues-raw/030/bcbsn.svg)
![Pull requests](https://img.shields.io/github/issues-pr-raw/030/bcbsn.svg)
![Total downloads](https://img.shields.io/github/downloads/030/bcbsn/total.svg)
![GitHub forks](https://img.shields.io/github/forks/030/bcbsn?label=fork&style=plastic)
![GitHub watchers](https://img.shields.io/github/watchers/030/bcbsn?style=plastic)
![GitHub stars](https://img.shields.io/github/stars/030/bcbsn?style=plastic)
![License](https://img.shields.io/github/license/030/bcbsn.svg)
![Repository Size](https://img.shields.io/github/repo-size/030/bcbsn.svg)
![Contributors](https://img.shields.io/github/contributors/030/bcbsn.svg)
![Commit activity](https://img.shields.io/github/commit-activity/m/030/bcbsn.svg)
![Last commit](https://img.shields.io/github/last-commit/030/bcbsn.svg)
![Release date](https://img.shields.io/github/release-date/030/bcbsn.svg)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/030/bcbsn?logo=github&sort=semver)](https://github.com/030/bcbsn/releases/latest)
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
[![codecov](https://codecov.io/gh/030/bcbsn/branch/main/graph/badge.svg)](https://codecov.io/gh/030/bcbsn)
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-web.svg)](https://golangci.com/r/github.com/030/bcbsn)
[![Chocolatey](https://img.shields.io/chocolatey/dt/bcbsn)](https://chocolatey.org/packages/bcbsn)
[![bcbsn](https://snapcraft.io//bcbsn/badge.svg)](https://snapcraft.io/bcbsn)
[![codebeat badge](https://codebeat.co/badges/f4aa5086-a4d5-41cd-893a-5da816ee9107)](https://codebeat.co/projects/github-com-030-bcbsn-main)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

[godoc]: https://godoc.org/github.com/030/bcbsn
[godoc widget]: https://godoc.org/github.com/030/bcbsn?status.svg

The Bitbucket Cloud Build Status Notifier (BCBSN) is a tool that sends a build
status to bitbucket cloud. Oauth credentials are required to perform this
action and these have to be created in the user settings menu that resides in
bitbucket cloud. Note: a Callback URL, e.g. <https://ci-system-name/> is
required in order to create Oauth credentials and one has to check 'private
consumer' if the build status has to be sent to private repositories.

In order to run this tool, one could choose to download the binary from the
releases section in this repository or use docker.

Example for running this command given:

- **client id**: 1234
- **client secret**: 4567
- **build state**: INPROGRESS
- **git commit hash**: f76d123498a053c1789057a41d6c3fcvg8b49cd7
- **owner**: 030
- **repository name**: bcbsn
- **build number**: 70
- **build url**: <https://travis-ci.org/030/bcbsn/builds/523263434/>

[![dockeri.co](https://dockeri.co/image/utrecht/bcbsn)](https://hub.docker.com/r/utrecht/bcbsn)

To print the help message

```bash
docker run utrecht/bcbsn:2.0.7 -h
```

returns:

```bash
Usage of /usr/local/bcbsn:
  -clientID string
      The clientID used for the 'client credentials' token flow with BitBucket
  -clientSecret string
      The clientSecret used for the 'client credentials' token flow with BitBucket
  -commit string
      The commit, e.g. 57484fd5460017aef111e8b4ec116a30ff0b4904
  -key string
      The key, e.g. a unique id of the build (use the build id)
  -name string
      An identifier for the build e.g. 'build 2'
  -owner string
      The owner of the repository, e.g. it is 'atlassian' in 'https://bitbucket.org/atlassian/stash-example-plugin/src/master/'
  -repoSlug string
      The repoSlug, e.g. some-repository
  -state string
      The state, e.g. SUCCESSFUL, INPROGRESS or FAILED
  -url string
      The url, e.g. https://travis-ci.org/030/bcbsn/builds/523263434
```

[godoc]: https://godoc.org/github.com/030/bcbsn
[godoc widget]: https://godoc.org/github.com/030/bcbsn?status.svg
