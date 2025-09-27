#!/bin/bash

echo "Starting Blockchain DApp Backend in Development Mode..."

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "Go is not installed. Please install Go first."
    exit 1
fi

# Check if PostgreSQL is running
if ! nc -z localhost 5432; then
    echo "PostgreSQL is not running on localhost:5432"
    echo "Please start PostgreSQL before running the backend."
    exit 1
fi

# Set environment variables
export DATABASE_URL="postgres://blockchain:blockchain@localhost:5432/blockchain_dev?sslmode=disable"
export PORT=8080

# Change to backend directory
cd backend

# Run the backend
echo "Starting backend server on port 8080..."
go run cmd/server/main.go