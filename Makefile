# Makefile for Simple Calculator (Go version)

# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GORUN = $(GOCMD) run
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test

# Build parameters
BINARY_NAME = calculator
BINARY_DIR = bin
MAIN_PATH = ./cmd/calculator

# Default target
all: build

# Build the executable
build:
	@mkdir -p $(BINARY_DIR)
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) $(MAIN_PATH)

# Run the calculator
run:
	$(GORUN) $(MAIN_PATH)

# Clean up build files
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_DIR)
	rm -f $(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Help target
help:
	@echo "Available targets:"
	@echo "  all     - Build the calculator (default)"
	@echo "  build   - Build the calculator binary"
	@echo "  run     - Run the calculator without building"
	@echo "  clean   - Remove build artifacts"
	@echo "  test    - Run all tests"
	@echo "  help    - Show this help message"

.PHONY: all build run clean test help
