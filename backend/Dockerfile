FROM golang:1.26.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY config.yaml .
EXPOSE 8080
VOLUME ["/data"]
CMD ["./main"]
