FROM golang:1.14.9-alpine
RUN mkdir /build
ADD go.mod go.sum hello.go /build/
WORKDIR /build
RUN go build