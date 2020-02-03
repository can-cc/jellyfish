FROM golang:1.12.1-stretch AS builder

WORKDIR /go/src/github.com/fwchen/jellyfish

COPY . .

ENV GO111MODULE=on

RUN GOPROXY='https://goproxy.cn' go get -v
RUN GOPROXY='https://goproxy.cn' go build cmd/jellyfish-server/main.go

FROM golang:1.12.1-stretch

WORKDIR /app

COPY --from=builder /go/src/github.com/fwchen/jellyfish/main /app
COPY --from=builder /go/src/github.com/fwchen/jellyfish/config /config

ENTRYPOINT ["/app/main"]
