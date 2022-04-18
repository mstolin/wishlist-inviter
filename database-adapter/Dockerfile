ARG BUILD_WORKDIR=/go/src/github.com/mstolin/present-roulette/database-adapter

FROM golang:1.18-rc-bullseye AS builder
LABEL stage=builder
COPY go.mod go.sum ${BUILD_WORKDIR}/
WORKDIR ${BUILD_WORKDIR}/
RUN go mod download
COPY . ${BUILD_WORKDIR}/
RUN GOOS=linux go build -a

FROM alpine:3.15.0
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=builder ${BUILD_WORKDIR}/database-adapter /usr/bin/database-adapter
EXPOSE 8080
ENTRYPOINT /usr/bin/database-adapter
