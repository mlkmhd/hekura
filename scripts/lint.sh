#!/bin/bash
# Check if golangci-lint is installed
if ! command -v golangci-lint &> /dev/null
then
    echo "golangci-lint could not be found, please install it: https://golangci-lint.run/usage/install/"
    exit 1
fi

golangci-lint run ./...
