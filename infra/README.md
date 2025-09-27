# Infrastructure Setup

This directory contains the Terraform configurations for setting up the AWS infrastructure.

## Account Setup

The AWS account (ID: 907849381252) has been configured with:

1. **CloudTrail** - For logging all AWS API calls
2. **Config** - For tracking configuration changes
3. **IAM Security** - With groups, policies, and strong password policies
4. **Secure Root Account** - With proper separation of duties

## VPC Topology & Network Plan

The infrastructure follows a well-designed VPC topology with:

- **Public Subnets** - For internet-facing resources
- **Private Subnets** - For application resources
- **Database Subnets** - For RDS instances
- **NAT Gateways** - For private subnet internet access
- **VPC Endpoints** - For private access to AWS services
- **VPC Flow Logs** - For network traffic monitoring

See [VPC Design Document](../docs/vpc-design.md) for detailed network architecture.

## Terraform Modules

Reusable Terraform modules are located in the [modules](modules/) directory:

- **[vpc](modules/vpc)** - Creates a complete VPC with public/private subnets, NAT gateways, and routing
- **[eks](modules/eks)** - Creates EKS clusters with node groups
- **[rds-postgres](modules/rds-postgres)** - Creates PostgreSQL database instances

## Environments

- **[dev](env/dev)** - Development environment
- **[staging](env/staging)** - Staging environment
- **[prod](env/prod)** - Production environment

## Global Resources

Global resources are defined in the [global](global/) directory:

- CloudTrail for API logging
- Config for resource tracking
- IAM policies and groups
- S3 buckets for Terraform state and application assets
- DynamoDB tables for state locking
- KMS keys for encryption
- ECR repositories for container images

## CI/CD

CI/CD workflows are defined in the [ci-cd](ci-cd/) directory:

- **[deploy-infra.yml](ci-cd/workflows/deploy-infra.yml)** - Infrastructure deployment
- **[deploy-app.yml](ci-cd/workflows/deploy-app.yml)** - Application deployment
- **[deploy-backend.yml](ci-cd/workflows/deploy-backend.yml)** - Backend deployment

See [CI/CD Setup Guide](../docs/ci-cd-setup.md) for detailed configuration instructions.

## Monitoring and Observability

The dev environment includes a complete monitoring stack:

- **Prometheus** - For metrics collection and alerting
- **Grafana** - For visualization and dashboarding
- **Alertmanager** - For handling alerts with Gmail SMTP or webhook notifications
- **dApp Health Checker** - Sample application with Prometheus metrics

See [Monitoring Setup Guide](../docs/monitoring-setup.md) for detailed configuration instructions.

## Setup Instructions

1. Ensure you have AWS credentials configured
2. Run `terraform init` in each environment directory
3. Run `terraform apply` to create the resources

## Security Features

- CloudTrail logs are stored in a secure S3 bucket with lifecycle policies
- Config records all resource configurations with delivery to S3
- IAM groups with least-privilege policies
- Strong password policy enforced
- Separate Terraform state storage per environment with versioning and encryption
- DynamoDB locking for Terraform state consistency
- Public access blocked on all S3 buckets
- VPC Flow Logs for network monitoring
- KMS encryption for data at rest
- IRSA (IAM Roles for Service Accounts) for secure Kubernetes access
- Secrets Manager for sensitive data