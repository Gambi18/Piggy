# Build stage
FROM golang:1.26-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY backend/ .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o piggy ./cmd/main.go

# Final stage
FROM alpine:3.19

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1001 -S piggy && \
    adduser -u 1001 -S piggy -G piggy

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/piggy .

# Copy migration files from builder stage
COPY --from=builder /app/internal/db/migrations ./internal/db/migrations

# Make binary executable
RUN chmod +x ./piggy

# Expose port
EXPOSE 8081

# Switch to non-root user
USER piggy

# Run the binary
CMD ["./piggy"]
