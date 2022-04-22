ARG WORKDIR=/go/src/github.com/mstolin/present-roulette

FROM golang:1.18-rc-bullseye AS builder
ARG SERVICE_NAME
ARG SERVICE_PROJECT_PATH
LABEL stage=builder
WORKDIR ${WORKDIR}/
COPY utils/ utils/
#COPY ${SERVICE_PROJECT_PATH}/go.mod ${SERVICE_PROJECT_PATH}/go.sum ${SERVICE_NAME}/
#WORKDIR ${SERVICE_NAME}/
#RUN go mod download
COPY ${SERVICE_PROJECT_PATH}/ ${SERVICE_NAME}/
RUN go work init utils/ ${SERVICE_NAME}/
RUN go work sync
WORKDIR ${SERVICE_NAME}/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./${SERVICE_NAME}

FROM alpine:3.15.0
ARG SERVICE_NAME
RUN apk --no-cache add ca-certificates
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=builder ${BUILD_WORKDIR}/${SERVICE_NAME}/${SERVICE_NAME} /usr/bin/${SERVICE_NAME}
EXPOSE 8080
#ENTRYPOINT /usr/bin/scrapper-facade
CMD /usr/bin/scrapper-facade
