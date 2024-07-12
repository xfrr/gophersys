# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install hot reload tool
RUN go install github.com/air-verse/air@latest

# Cache dependencies by copying go.mod and go.sum files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code to the working directory
COPY . .

ENTRYPOINT ["air", "-c", ".air.toml"]
