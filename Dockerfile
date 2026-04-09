# Stage 1: Build
FROM golang:1.25.8-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN go build -o main .

# Stage 2: Minimal runtime
FROM alpine:latest

WORKDIR /app

# (Optional but recommended) install CA certs
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .

EXPOSE 8085

CMD ["./main"]