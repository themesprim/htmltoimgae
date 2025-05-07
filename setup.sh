#!/bin/bash

# Exit on error
set -e

echo "Setting up HTML to Image Converter..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go first."
    exit 1
fi

# Check if wkhtmltoimage is installed
if ! command -v wkhtmltoimage &> /dev/null; then
    echo "Installing wkhtmltoimage..."
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        sudo apt-get update
        sudo apt-get install -y wkhtmltopdf
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        brew install wkhtmltopdf
    else
        echo "Please install wkhtmltopdf manually for your operating system"
        exit 1
    fi
fi

# Install Go dependencies
echo "Installing Go dependencies..."
go mod download

# Build the application
echo "Building the application..."
go build -o html-to-image

echo "Setup completed successfully!"
echo "You can now run the application with: ./html-to-image" 