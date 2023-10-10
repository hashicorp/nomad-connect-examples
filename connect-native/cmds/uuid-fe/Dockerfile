# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

FROM golang:alpine as builder
WORKDIR /build
ADD ./Dockerfile /build

RUN apk add git
RUN git clone https://github.com/hashicorp/nomad-connect-examples

WORKDIR /build/nomad-connect-examples/connect-native/cmds/uuid-fe
RUN go version && \
    go env && \
    CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3
MAINTAINER nomadproject.io

WORKDIR /opt
COPY --from=builder /build/nomad-connect-examples/connect-native/cmds/uuid-fe /opt

ENTRYPOINT ["/opt/uuid-fe"]
