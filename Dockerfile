FROM golang:1.12.1-alpine

ENV PROJECT /go/src/github.com/030/golang-bitbucket-cloud-build-status-notifier-linux/

RUN mkdir -p $PROJECT && \
    adduser -D -g '' gbcbsn

WORKDIR $PROJECT

COPY main.go Gopkg.toml ./

RUN apk add curl git && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep ensure && \
    CGO_ENABLED=0 go build && \
    cp $PROJECT/golang-bitbucket-cloud-build-status-notifier-linux /golang-bitbucket-cloud-build-status-notifier-linux

FROM scratch

COPY --from=0 /etc/passwd /etc/passwd
COPY --from=0 /golang-bitbucket-cloud-build-status-notifier-linux /usr/local/golang-bitbucket-cloud-build-status-notifier-linux
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER gbcbsn

ENTRYPOINT ["/usr/local/golang-bitbucket-cloud-build-status-notifier-linux"]
