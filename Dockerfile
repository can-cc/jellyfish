# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
ENTRYPOINT ./main