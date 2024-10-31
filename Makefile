# Name of the binary to be created
BINARY_NAME=birdcli

# Default Go build command
build:
	go build -o bin/$(BINARY_NAME) .

# Run the application
run:
	./bin/$(BINARY_NAME)

# Clean up the binary file
clean:
	rm -f bin/$(BINARY_NAME)

# Rebuild the application
rebuild: clean build

# Install the binary to /usr/local/bin
install: build
	sudo cp bin/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

# Delete the installed binary from /usr/local/bin
delete:
	sudo rm -f /usr/local/bin/$(BINARY_NAME)

# Help message to describe the make commands
help:
	@echo "Makefile commands:"
	@echo "  build   - Build the binary for birdcli"
	@echo "  run     - Run the birdcli binary"
	@echo "  clean   - Remove the birdcli binary from the bin folder"
	@echo "  rebuild - Clean and rebuild the birdcli binary"
	@echo "  install - Build and install birdcli to /usr/local/bin"
	@echo "  delete  - Remove the installed birdcli binary from /usr/local/bin"
	@echo "  help    - Show this help message"
