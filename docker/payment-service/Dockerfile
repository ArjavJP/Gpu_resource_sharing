# Use a Go 1.20 Alpine image as the base
FROM golang:1.23.0

# Install dependencies like git and protobuf if needed for building (optional)
RUN apt-get update && apt-get install -y git protobuf-compiler

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files to download dependencies first
COPY ../../go.mod ../../go.sum ./

# Copy the contents of the 'cmd/payment-service' folder to the container's /app directory
COPY ../../cmd/payment-service/ ./

# Build the Go application from the 'main.go' file
RUN go build -o payment-service .

# Expose the port your payment service will run on (e.g., port 5003)
EXPOSE 5003

# Command to run the payment service
CMD ["./payment-service"]
