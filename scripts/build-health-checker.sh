#!/bin/bash

# Script to build and push the dApp Health Checker image to ECR

set -e

# AWS configuration
AWS_ACCOUNT_ID="907849381252"
AWS_REGION="us-east-1"
ECR_REPOSITORY="dapp-health-checker"

# Get the ECR registry URI
ECR_REGISTRY="${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"

# Login to ECR
echo "Logging in to Amazon ECR..."
aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${ECR_REGISTRY}

# Build the Docker image
echo "Building Docker image..."
docker build -t ${ECR_REPOSITORY}:latest ../app/health-checker

# Tag the image for ECR
echo "Tagging image for ECR..."
docker tag ${ECR_REPOSITORY}:latest ${ECR_REGISTRY}/${ECR_REPOSITORY}:latest

# Push the image to ECR
echo "Pushing image to ECR..."
docker push ${ECR_REGISTRY}/${ECR_REPOSITORY}:latest

echo "Image successfully pushed to ECR!"
echo "Image URI: ${ECR_REGISTRY}/${ECR_REPOSITORY}:latest"