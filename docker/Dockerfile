# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache build-base

# Set necessary environment variables
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Cache dependencies by copying go.mod and go.sum files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o /app/main ./cmd/server

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Ensure the binary has the proper permissions
RUN chmod +x ./main

# Add a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Run the Go app by default when the container starts
CMD ["./main"]

# Optionally, add metadata to the image
LABEL maintainer="Francisco Romero <work@fromero.me>"
LABEL description="A simple CRUD application using Go and gRPC"

# grpc port
EXPOSE 50051

# http port
EXPOSE 8080