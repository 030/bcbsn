FROM golang:1.12.1-alpine

ENV PROJECT gbcbsn

RUN mkdir $PROJECT && \
    adduser -D -g '' $PROJECT

COPY main.go go.mod go.sum ./$PROJECT/

WORKDIR $PROJECT

RUN apk add git && \
    CGO_ENABLED=0 go build && \
    cp golang-bitbucket-cloud-build-status-notifier /golang-bitbucket-cloud-build-status-notifier

FROM scratch

COPY --from=0 /etc/passwd /etc/passwd
COPY --from=0 /golang-bitbucket-cloud-build-status-notifier /golang-bitbucket-cloud-build-status-notifier
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER $PROJECT

ENTRYPOINT ["/golang-bitbucket-cloud-build-status-notifier"]
