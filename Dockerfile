FROM golang:1.16 AS builder

RUN mkdir -p /go/src/github.com/hsmtkk/addhosts

WORKDIR /go/src/github.com/hsmtkk/addhosts

COPY getip .
COPY go.mod .
COPY go.sum .
COPY main.go .

ENV CGO_ENABLED=0

RUN go build -o addhosts.bin

FROM alpine:3.13.2

COPY --from=builder /opt/addhosts.bin /opt/addhosts.bin
