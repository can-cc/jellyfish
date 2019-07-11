FROM golang:1.12.1-stretch AS builder

WORKDIR /go/src/github/fwchen/jellyfish

COPY . .

ENV GO111MODULE=on

RUN go get -v

RUN go build main.go



FROM alpine

WORKDIR /app

COPY --from=builder /go/src/github.com/fwchen/jellyfish/main /app

EXPOSE 8180

CMD ["/app/main"]