FROM joosthofman/wget AS builder
RUN wget https://github.com/030/golang-bitbucket-cloud-build-status-notifier/releases/download/1.0.0/golang-bitbucket-cloud-build-status-notifier-linux --no-check-certificate && \
wget https://github.com/030/golang-bitbucket-cloud-build-status-notifier/releases/download/1.0.0/golang-bitbucket-cloud-build-status-notifier-linux.sha512.txt --no-check-certificate && \
sha512sum golang-bitbucket-cloud-build-status-notifier-linux && \
chmod +x golang-bitbucket-cloud-build-status-notifier-linux

FROM alpine
COPY --from=builder /golang-bitbucket-cloud-build-status-notifier-linux /golang-bitbucket-cloud-build-status-notifier-linux
RUN ls && mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ENTRYPOINT ["/golang-bitbucket-cloud-build-status-notifier-linux"]
