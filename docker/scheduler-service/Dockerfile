# Use a Go 1.20 Alpine image as the base
FROM golang:1.23.0
# Install dependencies like git and protobuf if needed for building (optional)
RUN apt-get update && apt-get install -y git protobuf-compiler

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files to download dependencies first
COPY ../../go.mod ../../go.sum ./


# Copy the contents of the 'cmd/scheduler-service' folder to the container's /app directory
COPY ../../cmd/scheduler-service/ ./

# Build the Go application from the 'main.go' file
RUN go build -o scheduler-service .

# Expose the port your scheduler service will run on (e.g., port 5002)
EXPOSE 5002

# Command to run the scheduler service
CMD ["./scheduler-service"]
