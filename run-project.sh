#!/bin/bash

echo "Starting Blockchain DApp..."

# Check if docker is installed
if ! command -v docker &> /dev/null
then
    echo "Docker is not installed. Please install Docker first."
    exit 1
fi

# Check if docker compose is installed
if command -v docker-compose &> /dev/null
then
    DOCKER_COMPOSE_CMD="docker-compose"
elif command -v docker &> /dev/null && docker compose version &> /dev/null
then
    DOCKER_COMPOSE_CMD="docker compose"
else
    echo "docker-compose is not installed. Please install docker-compose first."
    exit 1
fi

# Build and start all services
echo "Building and starting services..."
$DOCKER_COMPOSE_CMD up --build

echo "Blockchain DApp is now running!"
echo "Frontend: http://localhost:3000"
echo "Backend API: http://localhost:8080"
echo "Database: postgres://blockchain:blockchain@localhost:5432/blockchain_dev"