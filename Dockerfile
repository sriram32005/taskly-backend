# ---------- Stage 1: Build ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git (needed for some Go modules)
RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build a static binary (best for production)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# ---------- Stage 2: Runtime ----------
FROM alpine:latest

WORKDIR /app

# Install CA certs (important for HTTPS calls)
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/main .

# Copy .env if you want it inside container (optional)
# COPY .env .env

# Set production environment
ENV GIN_MODE=release

# Expose port (can be overridden by env)
EXPOSE 8085

# Run the app
CMD ["./main"]