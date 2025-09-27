#!/bin/bash

# Script to test the frontend build

set -e

echo "Testing frontend build..."

# Navigate to the app directory
cd /mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/app

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
  echo "Installing dependencies..."
  npm install
fi

# Run a basic build test
echo "Running build test..."
npm run build

# Check if build was successful
if [ -d "build" ]; then
  echo "Frontend build test successful!"
  echo "Build directory contains:"
  ls -la build
else
  echo "Frontend build test failed!"
  exit 1
fi

echo "Frontend is ready for deployment to web, iOS, and Android platforms."