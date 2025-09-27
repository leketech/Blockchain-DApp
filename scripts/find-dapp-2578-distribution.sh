#!/bin/bash

# Script to find the Dapp-2578 CloudFront distribution
# This script searches for CloudFront distributions that might be associated with the dapp-bucket-2578 S3 bucket

set -e

# Check if AWS CLI is installed
if ! command -v aws &> /dev/null; then
    echo "AWS CLI is not installed. Please install it first."
    exit 1
fi

echo "Searching for Dapp-2578 CloudFront distribution..."

# List all CloudFront distributions and look for the one tagged with Dapp-2578
aws cloudfront list-distributions \
    --query 'DistributionList.Items[*].{Id:Id,Comment:Comment,DomainName:DomainName}' \
    --output table

echo ""
echo "To find the exact Dapp-2578 distribution, look for entries with 'Dapp-2578' in the Comment field."
echo "The distribution is configured to use the S3 bucket named 'dapp-bucket-2578' as its origin."
echo "You can also use the following command to filter by tag (if tags are set):"
echo "aws cloudfront list-distributions --query 'DistributionList.Items[?Tags[?Key==\`Name\` && Value==\`Dapp-2578-cloudfront-dev\`]].{Id:Id,Comment:Comment,DomainName:DomainName}' --output table"