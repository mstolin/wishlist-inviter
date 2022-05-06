#
# BUILDER
#
FROM golang:1.18-buster AS builder
ARG WORKDIR=/app
ARG SERVICE_NAME
ARG SERVICE_PROJECT_PATH
LABEL stage=builder
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
# Call this http-service instead of SERVICE_NAME, because ARGS cannot be used in ENTRYPOINT or CMD
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/http-service

#
# RUNNER
#
FROM alpine:3.15.0
# Install certificates for HTTPS
RUN apk --no-cache add ca-certificates
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=builder /go/bin/http-service /bin/http-service
EXPOSE 8080
ENTRYPOINT /bin/http-service
