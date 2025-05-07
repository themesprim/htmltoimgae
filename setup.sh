#!/bin/bash

# Exit on error
set -e

echo "Setting up HTML to Image Converter..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go first."
    exit 1
fi

# Check if Chrome is installed
if ! command -v google-chrome &> /dev/null && ! command -v chromium &> /dev/null; then
    echo "Installing Chrome/Chromium..."
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # For Ubuntu/Debian
        if command -v apt-get &> /dev/null; then
            wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | sudo apt-key add -
            sudo sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list'
            sudo apt-get update
            sudo apt-get install -y google-chrome-stable
        # For CentOS/RHEL
        elif command -v yum &> /dev/null; then
            sudo yum install -y chromium
        # For Fedora
        elif command -v dnf &> /dev/null; then
            sudo dnf install -y chromium
        else
            echo "Please install Chrome/Chromium manually for your Linux distribution"
            exit 1
        fi
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        brew install --cask google-chrome
    else
        echo "Please install Chrome/Chromium manually for your operating system"
        exit 1
    fi
fi

# Install Go dependencies
echo "Installing Go dependencies..."
go mod download
go mod tidy

# Build the application
echo "Building the application..."
go build -o html-to-image

echo "Setup completed successfully!"
echo "You can now run the application with: ./html-to-image" 