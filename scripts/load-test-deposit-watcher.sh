#!/bin/bash

# Load test script for deposit watcher

echo "Starting load test for deposit watcher..."

# Set up test environment
echo "Setting up test environment..."
export TEST_DB_URL="postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable"

# Create test database
echo "Creating test database..."
createdb -U testuser -h localhost -p 5432 testdb || true

# Run database migrations
echo "Running database migrations..."
cd backend
go run cmd/migrate/main.go

# Start the deposit watcher service
echo "Starting deposit watcher service..."
go run cmd/deposit-watcher/main.go &
DEPOSIT_WATCHER_PID=$!

# Wait for service to start
sleep 5

# Generate test load
echo "Generating test load..."

# Simulate 100 deposit transactions
for i in {1..100}; do
    # Generate a random Ethereum address
    ADDRESS="0x$(openssl rand -hex 20)"
    
    # Insert test wallet into database
    psql -U testuser -h localhost -p 5432 -d testdb -c "INSERT INTO wallets (user_id, chain, address, public_key, balance, type, created_at, updated_at) VALUES (1, 'ethereum', '$ADDRESS', '$ADDRESS', 0, 'deposit', NOW(), NOW());"
    
    # Insert test transaction
    TX_HASH="0x$(openssl rand -hex 32)"
    AMOUNT=$(echo "scale=6; $RANDOM/1000000" | bc)
    psql -U testuser -h localhost -p 5432 -d testdb -c "INSERT INTO transactions (wallet_id, tx_hash, from_address, to_address, amount, chain, status, confirmations, fee, created_at, updated_at) VALUES (1, '$TX_HASH', '0x0000000000000000000000000000000000000000', '$ADDRESS', $AMOUNT, 'ethereum', 'confirmed', 12, 0.00021, NOW(), NOW());"
    
    echo "Generated test transaction $i: $TX_HASH for $AMOUNT ETH to $ADDRESS"
done

# Wait for processing
echo "Waiting for deposit watcher to process transactions..."
sleep 30

# Check results
echo "Checking results..."
RECONCILED_COUNT=$(psql -U testuser -h localhost -p 5432 -d testdb -t -c "SELECT COUNT(*) FROM transactions WHERE reconciled = true;" | xargs)
echo "Reconciled transactions: $RECONCILED_COUNT"

# Verify account balances
echo "Verifying account balances..."
psql -U testuser -h localhost -p 5432 -d testdb -c "SELECT * FROM accounts;"

# Stop the deposit watcher service
echo "Stopping deposit watcher service..."
kill $DEPOSIT_WATCHER_PID

echo "Load test completed."