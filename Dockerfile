FROM golang:1.14.2-alpine
ENV PROJECT gbcbsn
RUN mkdir $PROJECT && \
    adduser -D -g '' $PROJECT
COPY main.go go.mod go.sum ./$PROJECT/
WORKDIR $PROJECT
RUN apk add git && \
    CGO_ENABLED=0 go build && \
    cp bcbsn /bcbsn

FROM scratch
COPY --from=0 /etc/passwd /etc/passwd
COPY --from=0 /bcbsn /usr/local/bcbsn
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER $PROJECT
ENTRYPOINT ["/usr/local/bcbsn"]