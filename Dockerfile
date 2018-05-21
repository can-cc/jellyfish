# build stage
FROM golang:1.9.6-alpine3.7 AS build-env

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc make alpine-sdk libssh2 ca-certificates openssl

RUN mkdir /go/src/jellyfish
WORKDIR /go/src/jellyfish
COPY . .
RUN go get -d -v github.com/labstack/echo
RUN go get -d -v github.com/labstack/gommon/log
RUN go get -d -v github.com/mattn/go-sqlite3
CMD ["jellyfish"]
RUN go build main.go
RUN go build tools/create-user.go

# final stage
FROM alpine
RUN mkdir /jellyfish
COPY --from=build-env /go/src/jellyfish/ /jellyfish