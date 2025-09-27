#!/bin/bash
echo "🚀 Bootstrapping Blockchain-DApp..."

# Install deps
cd backend && go mod tidy
cd ../app && npm install
cd ../mobile && npm install

echo "✅ Done! Run 'make' for available commands."

chmod +x bootstrap.sh