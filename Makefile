# Makefile for Bitdefender plugin

# Variables
APP_NAME := bitdefender
BUILD_DIR := build
SRC_FILES := scan.go
DOCKER_IMAGE := bitdefender-plugin:latest
GO_CMD := go
LINT_CMD := golangci-lint
TEST_CMD := go test

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	$(GO_CMD) build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_FILES)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(TEST_CMD) ./...

# Lint the code
.PHONY: lint
lint:
	@echo "Running linter..."
	$(LINT_CMD) run

# Clean build files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Run the application in local mode
.PHONY: run
run:
	@echo "Running $(APP_NAME)..."
	$(GO_CMD) run $(SRC_FILES) web

# Build Docker image
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Run Docker container
.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run --rm -p 3993:3993 -v $(PWD)/malware:/malware $(DOCKER_IMAGE)

# Push Docker image (example)
.PHONY: docker-push
docker-push:
	@echo "Pushing Docker image to registry..."
	docker push $(DOCKER_IMAGE)

# Help target to display available commands
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build         - Build the binary"
	@echo "  test          - Run tests"
	@echo "  lint          - Lint the code"
	@echo "  clean         - Clean build files"
	@echo "  run           - Run the application locally"
	@echo "  docker-build  - Build the Docker image"
	@echo "  docker-run    - Run the application in Docker"
	@echo "  docker-push   - Push the Docker image to registry"

