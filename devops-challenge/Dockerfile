# Stage 1: Build Stage
FROM golang:1.15-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first for better layer caching
# This ensures go mod download only runs when dependencies change
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the application
# CGO_ENABLED=0 for static binary
# -ldflags="-w -s" to strip debug information and reduce binary size
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .

# Stage 2: Runtime Stage (using distroless for minimal size and better security)
FROM gcr.io/distroless/static-debian11

# Copy timezone data and certificates from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 9999

# Use non-root user (distroless nonroot user has UID 65532)
USER nonroot:nonroot

# Run the application
ENTRYPOINT ["/app/main"]
