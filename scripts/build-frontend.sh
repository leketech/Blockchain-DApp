#!/bin/bash

# Script to build the frontend for web, iOS, and Android

set -e

echo "Building frontend for web..."
cd /mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/app
npm run build

echo "Frontend built successfully for web!"

echo "To build for iOS and Android, you can use the following commands:"
echo "  For iOS: npm run ios"
echo "  For Android: npm run android"

echo "Note: You'll need to have Xcode (for iOS) and Android Studio (for Android) installed to build for mobile platforms."