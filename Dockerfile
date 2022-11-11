FROM golang:1.18.0-alpine3.14 as builder
ARG VERSION
ENV USERNAME bcbsn
RUN adduser -D -g '' $USERNAME
COPY . /go/${USERNAME}/
WORKDIR /go/${USERNAME}/cmd/${USERNAME}
RUN apk add --no-cache git=~2 && \
  CGO_ENABLED=0 go build -ldflags "-X main.Version=${VERSION}" -buildvcs=false && \
  cp bcbsn /bcbsn

FROM alpine:3.13.1
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bcbsn /usr/local/bcbsn
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER bcbsn
ENTRYPOINT ["/usr/local/bin/bcbsn"]
