#!/bin/bash

# Define the binary name and destination directory
BINARY_NAME="birdcli"
DEST_DIR="/usr/local/bin"

# Check if the binary already exists
if [ -f "$DEST_DIR/$BINARY_NAME" ]; then
    echo "$BINARY_NAME already exists in $DEST_DIR. Overwriting."
fi

# Copy the binary to the destination directory
sudo cp "$BINARY_NAME" "$DEST_DIR/"

# Make it executable
sudo chmod +x "$DEST_DIR/$BINARY_NAME"

echo "$BINARY_NAME installed successfully!"