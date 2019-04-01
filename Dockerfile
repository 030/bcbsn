FROM golang:1.12.1-alpine

ENV PROJECT /go/src/github.com/030/golang-bitbucket-cloud-build-status-notifier-linux

RUN mkdir -p $PROJECT && \
    adduser -D -g '' gbcbsn

WORKDIR $PROJECT

COPY main.go go.mod ./

RUN apk add git && \
    export GO111MODULE=on && \
    go mod download && \
    CGO_ENABLED=0 go build && \
    ls && \
    cp golang-bitbucket-cloud-build-status-notifier /golang-bitbucket-cloud-build-status-notifier

FROM scratch

COPY --from=0 /etc/passwd /etc/passwd
COPY --from=0 /golang-bitbucket-cloud-build-status-notifier /golang-bitbucket-cloud-build-status-notifier
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER gbcbsn

ENTRYPOINT ["/golang-bitbucket-cloud-build-status-notifier"]
