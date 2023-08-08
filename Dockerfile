# Stage 1: Build the Go application
FROM golang:1.17 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

# Stage 2: Create a minimal image for the Go application
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/myapp .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]