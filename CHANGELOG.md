# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [2.0.2] - 2019-05-21
### Fixed
- Accept either a 200 or 201 as a HTTP response status code [@jorianvo]

## [2.0.1] - 2019-05-14
### Fixed
- List of arguments required to use this tool updated.
- Publication of artifacts.

## [2.0.0] - 2019-05-12
### Added
- HTTP tests [@jorianvo](https://github.com/jorianvo).
- Multiple badges including devops.stackexchange batch about [bcbsn tag](https://devops.stackexchange.com/questions/tagged/bcbsn).
- New dockerhub repository [utrecht/bcbsn](https://cloud.docker.com/u/utrecht/repository/docker/utrecht/bcbsn) as it was
impossible to rename repositories on dockerhub.

### Changed
- Repository renamed to bcbsn as former name exceeded tag character limit on [StackExchange](https://stackexchange.com/).
- All references in README renamed due to repository name change.

## [1.1.2] - 2019-05-06
### Added
- Documented how to use the bcbsn tool.

## [1.1.1] - 2019-04-23
### Fixed
- Executable file in Docker image was incorrect.

## [1.1.0] - 2019-04-23
### Changed
- Improved logging to ensure that it will become easier to use the bcbsn tool.

## [1.0.2] - 2019-03-26
### Fixed
- Access token missing due to missing CA certificates in Scratch image.

## [1.0.1] - 2019-03-17
### Changed
- Use Scratch as a base to get a smaller docker image.

## [1.0.0] - 2019-03-17
### Added
- Send a build status to bitbucket cloud.

[Unreleased]: https://github.com/030/bcbsn/compare/2.0.2...HEAD
[2.0.2]: https://github.com/030/bcbsn/compare/2.0.1...2.0.2
[2.0.1]: https://github.com/030/bcbsn/compare/2.0.0...2.0.1
[2.0.0]: https://github.com/030/bcbsn/compare/1.1.2...2.0.0
[1.1.2]: https://github.com/030/bcbsn/compare/1.1.1...1.1.2
[1.1.1]: https://github.com/030/bcbsn/compare/1.1.0...1.1.1
[1.1.0]: https://github.com/030/bcbsn/compare/1.0.2...1.1.0
[1.0.2]: https://github.com/030/bcbsn/compare/1.0.1...1.0.2
[1.0.1]: https://github.com/030/bcbsn/compare/1.0.0...1.0.1
[1.0.0]: https://github.com/030/bcbsn/releases/tag/1.0.0
