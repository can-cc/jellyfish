FROM golang:1.12.1-stretch

WORKDIR /go/src/github/fwchen/jellyfish

COPY . .

ENV GO111MODULE=on

RUN go get -v

RUN go build main.golang

CMD ["/go/src/github.com/fwchen/jellyfish/main"]