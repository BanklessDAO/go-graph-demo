ARG ALPINE_VERSION
FROM golang:alpine${ALPINE_VERSION} AS builder
WORKDIR /build
COPY . .
RUN \
    apk add -U --no-cache ca-certificates tzdata \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app . \
    && echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc_passwd /etc/passwd
COPY --from=builder /build/app /app
USER nobody
EXPOSE 8080
ENTRYPOINT [ "/app" ]