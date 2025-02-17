# Project variables
BINARY_NAME=image2pdf

# Build the application
build:
	go build -o $(BINARY_NAME) .

# Build for Windows PowerShell
ps-build:
	go build -o $(BINARY_NAME).exe .

# Run the application
run: build
	./$(BINARY_NAME)

# Run tests
test:
	go test ./...

# Format the code
fmt:
	go fmt ./...

# Clean up build files
clean:
	@if [ -e "$(BINARY_NAME)" ] || [ -e "$(BINARY_NAME).exe" ]; then \
		rm -f $(BINARY_NAME) $(BINARY_NAME).exe; \
	else \
		echo "No binaries to clean."; \
	fi

# Windows compatibility
clean-win:
	if exist $(BINARY_NAME).exe del /F /Q $(BINARY_NAME).exe
	if exist $(BINARY_NAME) del /F /Q $(BINARY_NAME)

# Tidy dependencies
tidy:
	go mod tidy

# Lint the code (requires golangci-lint)
lint:
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint not installed. Run 'go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest'"; exit 1; }
	golangci-lint run


# Install dependencies
install:
	go mod download

.PHONY: build ps-build run test fmt clean clean-win tidy lint install
