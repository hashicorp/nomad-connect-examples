FROM golang:alpine as builder
WORKDIR /build

RUN apk add git
RUN clone git@github.com:hashicorp/nomad-connect-examples.git

WORKDIR /build/nomad-connect-examples/connect-native/cmds/uuid-api
RUN go version && \
    go env && \
    CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3
MAINTAINER nomadproject.io

WORKDIR /opt
COPY --from=builder /build/nomad-connect-examples/connect-native/cmds/uuid-api /opt

ENTRYPOINT ["/opt/uuid-api"]
