FROM golang:1.15.3-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum /app/line_json.go /build/
WORKDIR /build
RUN go build line_json.go

FROM builder
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/line_json /app/helloworld
COPY views/ /app/views
WORKDIR /app

EXPOSE 8080
CMD ["./helloworld"]