# Build stage
FROM golang:1.22.3-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
# Targeting the main.go in cmd directory
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/binary ./cmd/main.go

# Final stage
FROM alpine:3.19

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/binary /app/binary

# Copy any necessary config files if needed
COPY --from=builder /app/config /app/config

# Create a non-root user
RUN adduser -D appuser
USER appuser

# Expose the port your application uses
EXPOSE 800

# Run the binary
CMD ["/app/binary"]