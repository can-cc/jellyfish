FROM golang:1.12.1-stretch AS builder

WORKDIR /go/src/github.com/fwchen/jellyfish

COPY . .

ENV GO111MODULE=on

RUN go get -v

RUN go build main.go

RUN ls


FROM golang:1.12.1-stretch

WORKDIR /app

COPY --from=builder /go/src/github.com/fwchen/jellyfish/main /app
COPY --from=builder /go/src/github.com/fwchen/jellyfish/config.yaml /app

ENTRYPOINT ["/app/main"]
