# Name of the binary to be created
BINARY_NAME=birdcli

# Default Go build command
build:
	go build -o $(BINARY_NAME) .

# Run the application
run:
	./$(BINARY_NAME)

# Clean up the binary file
clean:
	rm -f $(BINARY_NAME)

# Rebuild the application
rebuild: clean build

# Install the binary to /usr/local/bin
install: build
	./install.sh

# Help message to describe the make commands
help:
	@echo "Makefile commands:"
	@echo "  build   - Build the binary for birdcli"
	@echo "  run     - Run the birdcli binary"
	@echo "  clean   - Remove the birdcli binary"
	@echo "  rebuild - Clean and rebuild the birdcli binary"
	@echo "  install - Build and install birdcli to /usr/local/bin"
	@echo "  help    - Show this help message"
