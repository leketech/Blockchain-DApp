# VPC Topology & Network Plan

## Overview

This document describes the VPC topology and network design for the Blockchain-DApp across all environments (dev, staging, prod).

## Network Design

### CIDR Block Allocation

- **VPC CIDR**: 10.0.0.0/16 (65,536 IPs)
- **Availability Zones**: 3 AZs (us-east-1a, us-east-1b, us-east-1c)

### Subnet Design

#### Public Subnets (for internet-facing resources)
- **Public Subnet 1**: 10.0.1.0/24 (256 IPs) - us-east-1a
- **Public Subnet 2**: 10.0.2.0/24 (256 IPs) - us-east-1b
- **Public Subnet 3**: 10.0.3.0/24 (256 IPs) - us-east-1c

#### Private Subnets (for application resources)
- **Private Subnet 1**: 10.0.11.0/24 (256 IPs) - us-east-1a
- **Private Subnet 2**: 10.0.12.0/24 (256 IPs) - us-east-1b
- **Private Subnet 3**: 10.0.13.0/24 (256 IPs) - us-east-1c

#### Database Subnets (for RDS)
- **Database Subnet 1**: 10.0.21.0/24 (256 IPs) - us-east-1a
- **Database Subnet 2**: 10.0.22.0/24 (256 IPs) - us-east-1b
- **Database Subnet 3**: 10.0.23.0/24 (256 IPs) - us-east-1c

### Network Components

1. **Internet Gateway**: Attached to VPC for public subnet internet access
2. **NAT Gateways**: One per AZ in public subnets for private subnet internet access
3. **Route Tables**:
   - Public route table (routes to IGW)
   - Private route tables (routes to respective NAT gateways)
   - Database route table (no internet access)
4. **Security Groups**:
   - Web/Application tier SG
   - Database tier SG
   - Bastion host SG (if needed)
5. **Network ACLs**: Default VPC NACLs (can be customized if needed)

### VPC Endpoints

- **S3 Gateway Endpoint**: For private access to S3
- **DynamoDB Gateway Endpoint**: For private access to DynamoDB

## Environment-Specific Configurations

### Dev Environment
- Smaller instance types
- Single NAT Gateway to reduce costs
- Minimal node group configuration

### Staging Environment
- Medium instance types
- NAT Gateway per AZ
- Standard node group configuration

### Prod Environment
- Larger instance types
- NAT Gateway per AZ
- Highly available node group configuration
- Additional monitoring and alerting

## Security Considerations

1. **Network Segmentation**: Resources separated by subnet tiers
2. **Least Privilege**: Security groups with minimal required access
3. **Encryption**: All data in transit encrypted with TLS
4. **Logging**: VPC Flow Logs enabled for network traffic monitoring