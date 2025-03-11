# Makefile to build and run services

.PHONY: build all run clean

# Build all services
all: build

# Build the network, payment, and scheduler services
build: build-network-service build-payment-service build-scheduler-service

build-network-service:
	cd cmd/network-service && go build -o network-service

build-payment-service:
	cd cmd/payment-service && go build -o payment-service

build-scheduler-service:
	cd cmd/scheduler-service && go build -o scheduler-service

# Run all services with Docker Compose
run:
	docker-compose up --build

# Clean build artifacts
clean:
	rm -rf cmd/network-service/network-service
	rm -rf cmd/payment-service/payment-service
	rm -rf cmd/scheduler-service/scheduler-service
