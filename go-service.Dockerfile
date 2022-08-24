#
# BUILDER
#
FROM docker.io/golang:1.19-alpine AS builder
ARG WORKDIR=$GOPATH/src/github.com/mstolin/wishlist-inviter/
ARG SERVICE_NAME
ARG SERVICE_PROJECT_PATH
LABEL stage=builder
# lib64
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# Copy utils, this is needed in all modules
ADD utils/ ${WORKDIR}/utils
# Copy the desired module
ADD ${SERVICE_PROJECT_PATH}/ ${WORKDIR}/${SERVICE_NAME}
# Create go workspace and sync dependencies
WORKDIR ${WORKDIR}/
RUN go work init ./utils ./${SERVICE_NAME}
RUN go work sync
# Build service binary
WORKDIR ${WORKDIR}/${SERVICE_NAME}
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/http-service

#
# RUNNER
#
FROM alpine:3.15.0
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=builder /go/bin/http-service /go/bin/http-service
EXPOSE 8080
ENTRYPOINT ["/go/bin/http-service"]
