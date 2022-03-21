FROM golang:alpine as build
WORKDIR /build
RUN apk add --no-cache git ca-certificates build-base su-exec olm-dev
COPY . .
RUN go build -o standupbot

FROM alpine:latest
RUN apk add --no-cache ca-certificates olm-dev sqlite
WORKDIR /opt/standupbot
ENV PATH="${PATH}:/opt/standupbot"
COPY --from=build /build/standupbot ./
STOPSIGNAL SIGTERM
ENTRYPOINT ["standupbot"]

