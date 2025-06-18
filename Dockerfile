# Build stage
FROM golang:1.24.3-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o loadbalancer main.go

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/loadbalancer .
COPY config/config.json . 
EXPOSE 8080
CMD ["./loadbalancer"]