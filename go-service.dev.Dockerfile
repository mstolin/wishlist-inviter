FROM golang:1.18-buster AS builder
ARG WORKDIR=/app
ARG SERVICE_PROJECT_PATH
LABEL stage=builder
# Copy utils, this is needed in all modules
ADD utils/ ${WORKDIR}/utils
# Copy the desired module
ADD ${SERVICE_PROJECT_PATH}/ ${WORKDIR}/http-service
# Create go workspace and sync dependencies
WORKDIR ${WORKDIR}/
RUN go work init ./utils ./http-service
RUN go work sync
# Expose port
EXPOSE 8080
# Start service
ENTRYPOINT go run ./http-service
