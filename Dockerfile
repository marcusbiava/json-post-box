# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o json-post-box ./cmd/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/json-post-box .

EXPOSE 8030

CMD ["./json-post-box"]