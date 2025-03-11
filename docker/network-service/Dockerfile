# Use a Go 1.20 Alpine image as the base
FROM golang:1.23.0
# Install dependencies like git and protobuf if needed for building (optional)
RUN apt-get update && apt-get install -y git protobuf-compiler

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files to download dependencies first
COPY ../../go.mod ../../go.sum ./


# Copy the contents of the 'cmd/network-service' folder to the container's /app directory
COPY ../../cmd/network-service/ ./

# Build the Go application from the 'main.go' file
RUN go build -o network-service .

# Expose the port your network service will run on (e.g., port 5001)
EXPOSE 5001

# Command to run the network service
CMD ["./network-service"]
