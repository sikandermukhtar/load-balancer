FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o load-balancer main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/load-balancer .
CMD ["./load-balancer"]
