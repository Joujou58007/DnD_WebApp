# Build stage (common for dev and prod)
FROM golang:1.24.1 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Development stage
FROM golang:1.24.1 AS dev
WORKDIR /app
COPY . .
RUN go install github.com/air-verse/air@latest  # Updated to the correct path
EXPOSE 8080
CMD ["air"]

# Production stage
FROM alpine:latest AS prod
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]