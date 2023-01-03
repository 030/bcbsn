FROM golang:1.19.4-alpine3.17 as builder
ARG VERSION
ENV USERNAME bcbsn
RUN adduser -D -g '' $USERNAME
COPY . /go/${USERNAME}/
WORKDIR /go/${USERNAME}/cmd/${USERNAME}
RUN apk add --no-cache git=~2 && \
  CGO_ENABLED=0 go build -ldflags "-X main.Version=${VERSION}" -buildvcs=false && \
  cp bcbsn /bcbsn

FROM alpine:3.17.0
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bcbsn /usr/local/bin/bcbsn
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER bcbsn
ENTRYPOINT ["/usr/local/bin/bcbsn"]
