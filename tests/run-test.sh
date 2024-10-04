#!/bin/bash

# Stop on first error
set -e

# Initialize the Go module
echo "Initializing Go module..."
go mod init github.com/eazycloudlife/terraform-azure-labels

# Tidy the Go modules
echo "Tidying up Go modules..."
go mod tidy

# Run the Go tests with a timeout of 30 minutes
echo "Running Go tests with a timeout of 30 minutes..."
go test -v -timeout 30m

# Check if the test was successful
if [ $? -eq 0 ]; then
    echo "Test passed successfully."
else
    echo "Test failed."
    exit 1
fi
