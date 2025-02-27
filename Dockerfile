# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application with optimizations
# Note: Building from cmd/server where main.go is located
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Runtime stage
FROM alpine:latest

# Add CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create a non-root user to run the application
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy necessary directories
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/data ./data
COPY --from=builder /app/.env ./

# Expose the port the app runs on
EXPOSE 8080

# Switch to non-root user
USER appuser

# Command to run the application
CMD ["./main"]