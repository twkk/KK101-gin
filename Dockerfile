FROM golang:1.14.9-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum bank1.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/helloworld /app/
COPY views/ /app/views
WORKDIR /app
CMD ["./helloworld"]