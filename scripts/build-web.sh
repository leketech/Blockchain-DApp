#!/bin/bash

# Script to build the web version of the dApp

set -e

echo "Building web version of dApp..."

# Navigate to the web directory
cd /mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/web

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
  echo "Installing dependencies..."
  npm install
fi

# Run a basic build
echo "Running build..."
npm run build

# Check if build was successful
if [ -d "dist" ]; then
  echo "Web build successful!"
  echo "Build directory contains:"
  ls -la dist
else
  echo "Web build failed!"
  exit 1
fi

echo "Web version is ready for deployment."