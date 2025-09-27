# Use the official Golang image as the base image
FROM golang:1.24-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Use a minimal alpine image for the final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN adduser -D -s /bin/sh appuser

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Change ownership of the binary to the non-root user
RUN chown appuser:appuser main

# Switch to the non-root user
USER appuser

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]