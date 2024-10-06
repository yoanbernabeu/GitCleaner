#!/bin/bash

# Set variables
REPO_NAME="yoanbernabeu/GitCleaner"
BINARY_NAME="git-cleaner"
VERSION="0.1.3"

# Determine the OS and architecture
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$OS" == "darwin" && "$ARCH" == "arm64" ]]; then
    ARCH="arm64" # Apple M1 chip
elif [[ "$ARCH" == "x86_64" ]]; then
    ARCH="amd64"
elif [[ "$ARCH" == "aarch64" || "$ARCH" == "armv8l" ]]; then
    ARCH="arm64"
elif [[ "$ARCH" == "i386" || "$ARCH" == "i686" ]]; then
    ARCH="386"
fi

# Delete previous installation
if [ -f /usr/local/bin/$BINARY_NAME ]; then
    echo "Removing previous installation..."
    sudo rm /usr/local/bin/$BINARY_NAME
fi

# Construct the download URL
DOWNLOAD_URL="https://github.com/$REPO_NAME/releases/download/$VERSION/$BINARY_NAME-$VERSION-$OS-$ARCH.tar.gz"

# Download the archive
echo "Downloading $BINARY_NAME from $DOWNLOAD_URL..."
curl -L $DOWNLOAD_URL -o $BINARY_NAME.tar.gz

# Extract the archive
echo "Extracting $BINARY_NAME..."
tar -xzf $BINARY_NAME.tar.gz

# Remove the archive
rm $BINARY_NAME.tar.gz

# Check if the binary exists
if [ ! -f $BINARY_NAME ]; then
    echo "Error: Binary not found. Please try again."
    exit 1
fi

# Add the binary to the PATH
echo "Adding $BINARY_NAME to PATH..."
sudo mv $BINARY_NAME /usr/local/bin

# Check if the binary was added successfully
if [ ! -f /usr/local/bin/$BINARY_NAME ]; then
    echo "Error: Binary not added to PATH. Please try again."
    exit 1
fi

# Print success message
echo "Installation complete! $BINARY_NAME is now available in your PATH."