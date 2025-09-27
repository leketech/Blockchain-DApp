#!/bin/bash

# Script to test the web build

set -e

echo "Testing web build..."

# Navigate to the web directory
cd /mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/web

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
  echo "Installing dependencies..."
  npm install
fi

# Run a basic build test
echo "Running build test..."
npm run build

# Check if build was successful
if [ -d "dist" ]; then
  echo "Web build test successful!"
  echo "Build directory contains:"
  ls -la dist
else
  echo "Web build test failed!"
  exit 1
fi

echo "Web interface is ready for deployment."