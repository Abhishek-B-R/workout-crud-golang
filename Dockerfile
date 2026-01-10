# --- Stage 1: Build ---
# Use a specific version of the official Go image as the build stage base
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container for the build stage
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker's build cache
COPY go.mod go.sum ./

# Download all dependencies. This layer is cached efficiently if the mod files haven't changed.
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application into a static binary
# CGO_ENABLED=0 ensures a statically linked binary, which is portable
# -o /app/server names the output binary 'server' and places it in /app
RUN CGO_ENABLED=0 go build -o /app/server .

# --- Stage 2: Runtime ---
# Start a new, minimal stage from scratch (or alpine for basic tooling)
FROM alpine:latest

# Set the working directory for the final stage
WORKDIR /app

# Copy only the compiled binary from the 'builder' stage
COPY --from=builder /app/server .

# Expose the port your application listens on (e.g., 8080)
EXPOSE 8080

# Command to run the executable when the container starts
ENTRYPOINT ["/app/server"]
