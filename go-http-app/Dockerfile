# Multi-stage build to keep final image tiny and reduce vulnerabilities

# Builder stage (must be >= Go 1.24 because go.mod requires it)
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build a statically-linked binary for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./

# Final runtime image (minimal, no shell, very small attack surface)
FROM scratch

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/server /app/server

# Expose port your app listens on
EXPOSE 8080

# Run as non-root user by default UID:GID 65532 (nobody)
USER 65532:65532

ENTRYPOINT ["/app/server"]
