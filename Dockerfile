FROM golang:1.12.1-stretch AS builder

WORKDIR /go/src/github.com/fwchen/jellyfish

COPY . .

ENV GO111MODULE=on

RUN GOPROXY='https://goproxy.cn' go get -v
RUN GOPROXY='https://goproxy.cn' go build cmd/jellyfish-server/main.go

ENTRYPOINT ["/go/src/github.com/fwchen/jellyfish/main"]
