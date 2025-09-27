#!/bin/bash

# Script to initialize the mobile project

set -e

echo "Initializing mobile project..."

# Navigate to the mobile directory
cd /mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/mobile

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
  echo "Installing dependencies..."
  npm install
fi

# Check if Expo CLI is installed
if ! command -v expo &> /dev/null
then
  echo "Installing Expo CLI..."
  npm install -g expo-cli
fi

echo "Mobile project initialized successfully!"
echo "To start the development server, run: npm start"
echo "To run on iOS, run: npm run ios"
echo "To run on Android, run: npm run android"