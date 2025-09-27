#!/bin/bash

echo "Setting up PostgreSQL database for Blockchain DApp..."

# Check if PostgreSQL is installed
if ! command -v psql &> /dev/null
then
    echo "PostgreSQL is not installed. Please install PostgreSQL first."
    exit 1
fi

# Check if PostgreSQL is running
if ! nc -z localhost 5432; then
    echo "PostgreSQL is not running on localhost:5432"
    echo "Please start PostgreSQL before running this script."
    exit 1
fi

# Create database and user
echo "Creating database and user..."
psql -U postgres -c "CREATE USER blockchain WITH PASSWORD 'blockchain';" 2>/dev/null || echo "User already exists"
psql -U postgres -c "CREATE DATABASE blockchain_dev OWNER blockchain;" 2>/dev/null || echo "Database already exists"
psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE blockchain_dev TO blockchain;" 2>/dev/null || echo "Privileges already granted"

echo "Database setup completed!"
echo "Connection string: postgres://blockchain:blockchain@localhost:5432/blockchain_dev"