# Certs
FROM alpine:latest as certs
RUN apk --update add ca-certificates

# Build
FROM golang:1.13.4-alpine as base
WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN go build -o app -ldflags '-s -w'

# App
FROM scratch as app
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=base app /
EXPOSE 8080
ENTRYPOINT ["/app"]