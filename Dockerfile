FROM golang:1.17.0-alpine3.14 as builder
ENV PROJECT gbcbsn
RUN mkdir $PROJECT && \
    adduser -D -g '' $PROJECT
COPY main.go go.mod go.sum ./$PROJECT/
WORKDIR $PROJECT
RUN apk add git && \
    CGO_ENABLED=0 go build && \
    cp bcbsn /bcbsn

FROM alpine:3.13.1
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bcbsn /usr/local/bcbsn
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER $PROJECT
ENTRYPOINT ["/usr/local/bcbsn"]
