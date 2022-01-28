FROM golang:1.17-stretch AS builder

WORKDIR /app

COPY . .

ENV GO111MODULE=on

RUN GOPROXY='https://goproxy.cn' go get -v ./cmd/server
RUN GOPROXY='https://goproxy.cn' go build ./cmd/server/main.go

FROM golang:1.17-stretch

WORKDIR /app

COPY --from=builder /app/main /app
RUN mkdir /app/config
COPY ./config/config.yaml /app/config/config.yaml

ENTRYPOINT ["/app/main"]
