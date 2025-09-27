# Scripts

This directory contains various scripts for building, testing, and deploying the Blockchain DApp.

## Build Scripts

- [build-frontend.sh](file:///mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/scripts/build-frontend.sh) - Build the frontend application
- [build-health-checker.sh](file:///mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/scripts/build-health-checker.sh) - Build the health checker service
- [build-web.sh](file:///mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/scripts/build-web.sh) - Build the web application

## Test Scripts

- [test-frontend.sh](file:///mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/scripts/test-frontend.sh) - Run frontend tests
- [test-mobile.sh](file:///mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/scripts/test-mobile.sh) - Run mobile app tests
- [test-web.sh](file:///mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/scripts/test-web.sh) - Run web application tests

## Deployment Scripts

- [deploy-monitoring.sh](file:///mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/scripts/deploy-monitoring.sh) - Deploy monitoring stack
- [init-mobile.sh](file:///mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/scripts/init-mobile.sh) - Initialize mobile application
- [load-test-deposit-watcher.sh](file:///mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/scripts/load-test-deposit-watcher.sh) - Run load test on deposit watcher

## CloudFront Management Scripts

- [invalidate-cloudfront.sh](file:///mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/scripts/invalidate-cloudfront.sh) - Invalidate CloudFront cache for Dapp-2578 distribution
- [find-dapp-2578-distribution.sh](file:///mnt/c/Users/Leke/Decentralized-App/Blockchain-DApp/scripts/find-dapp-2578-distribution.sh) - Find the Dapp-2578 CloudFront distribution

## Usage

Make sure to make the scripts executable before running them:
```bash
chmod +x scripts/*.sh
```

Then run any script with:
```bash
./scripts/script-name.sh
```