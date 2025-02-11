# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

WORKDIR /app

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o goapp

# Stage 2: Create a lightweight runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/goapp .

# Command to run the application
CMD ["./goapp"]
