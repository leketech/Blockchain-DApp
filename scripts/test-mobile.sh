#!/bin/bash

# Script to test the mobile build

set -e

echo "Testing mobile build..."

# Navigate to the mobile directory
cd /mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/mobile

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
  echo "Installing dependencies..."
  npm install
fi

# Run a basic build test
echo "Checking if project compiles..."
npx tsc --noEmit

echo "Mobile build test completed!"
echo "To run the app, use: npm start"