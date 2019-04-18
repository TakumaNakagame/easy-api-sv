FROM golang:1.11.9-alpine3.9 as builder
MAINTAINER kameneko
ENV GOARCH="amd64"
ENV GOOS="linux"

COPY . /go/src/
WORKDIR /go/src
RUN apk update && apk add git
RUN /usr/local/go/bin/go get github.com/labstack/echo github.com/dgrijalva/jwt-go &&\
    /usr/local/go/bin/go build -o easy-api-sv main.go

FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/easy-api-sv /easy-api-sv
EXPOSE 8080
CMD ["/easy-api-sv"]