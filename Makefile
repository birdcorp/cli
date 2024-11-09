# Name of the binary to be created
BINARY_NAME=birdcli

# Default Go build command
build:
	go build -o bin/$(BINARY_NAME) .

# Run the application
run:
	./bin/$(BINARY_NAME)

# Clean up the binary file from the bin folder
clean:
	rm -f bin/$(BINARY_NAME)

# Rebuild the application
rebuild: clean build

# Install the binary to /usr/local/bin
install: build
	sudo cp bin/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

# Uninstall the binary from /usr/local/bin
uninstall:
	sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "$(BINARY_NAME) has been uninstalled from /usr/local/bin"

# Help message to describe the make commands
help:
	@echo "Makefile commands:"
	@echo "  build     - Build the binary for $(BINARY_NAME)"
	@echo "  run       - Run the $(BINARY_NAME) binary"
	@echo "  clean     - Remove the $(BINARY_NAME) binary from the bin folder"
	@echo "  rebuild   - Clean and rebuild the $(BINARY_NAME) binary"
	@echo "  install   - Build and install $(BINARY_NAME) to /usr/local/bin"
	@echo "  uninstall - Remove the installed $(BINARY_NAME) binary from /usr/local/bin"
	@echo "  help      - Show this help message"
