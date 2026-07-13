FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Disable Cgo to produce a static binary with no C dependencies (e.g., SQLite, glibc).
# Ensures portability across Linux hosts and minimal final image size.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY config.yaml .
EXPOSE 8080
VOLUME ["/data"]
CMD ["./main"]
