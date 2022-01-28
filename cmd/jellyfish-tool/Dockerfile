FROM golang:1.12.1-stretch AS builder

WORKDIR /go/src/jellyfish

COPY . .

ENV GO111MODULE=on

RUN GOPROXY='https://goproxy.cn' go build -o tool cmd/jellyfish-tool/main.go

FROM golang:1.12.1-stretch

WORKDIR /app

COPY --from=builder /go/src/jellyfish/tool /app

ENTRYPOINT ["/app/tool"]
