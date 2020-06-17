FROM golang:alpine as builder
WORKDIR /build

RUN apk add git && \
    git clone https://github.com/hashicorp/nomad-connect-examples

WORKDIR /build/nomad-connect-examples/cmds/uuid-fe
RUN go version && \
    go env && \
    CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3
MAINTAINER nomadproject.io

WORKDIR /opt
COPY --from=builder /build/connect-native-demo/cmds/uuid-fe /opt

ENTRYPOINT ["/opt/uuid-fe"]
