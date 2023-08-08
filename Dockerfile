# Set the base image to use for your Go application. We'll use the official Go image.
FROM golang:1.20

# Set the working directory inside the container.
WORKDIR /app

# Copy the Go application source code into the container's working directory.
COPY . .

# Build the Go application inside the container.
RUN go build -o app

# Expose any required ports. (This depends on your application)
EXPOSE 80

# Set the command to run your Go application when the container starts.
CMD ["./app"]
