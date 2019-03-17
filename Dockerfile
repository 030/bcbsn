FROM golang:alpine AS builder

ENV PROJECT /go/src/github.com/030/golang-bitbucket-cloud-build-status-notifier-linux/

RUN mkdir -p $PROJECT

WORKDIR $PROJECT

COPY main.go Gopkg.toml ./
RUN apk add curl git && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep ensure && \
    CGO_ENABLED=0 go build && \
    cp $PROJECT/golang-bitbucket-cloud-build-status-notifier-linux /golang-bitbucket-cloud-build-status-notifier-linux

FROM scratch
COPY --from=builder /golang-bitbucket-cloud-build-status-notifier-linux /golang-bitbucket-cloud-build-status-notifier-linux
ENTRYPOINT ["/golang-bitbucket-cloud-build-status-notifier-linux"]
