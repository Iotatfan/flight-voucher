FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build the application
COPY . .
# Disable CGO for a statically linked binary and build
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a minimal alpine image for the final stage
FROM alpine:latest
WORKDIR /app

# Copy the built binary and config from the builder stage
COPY --from=builder /app/main .
COPY config.yaml .

# Expose the server port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
