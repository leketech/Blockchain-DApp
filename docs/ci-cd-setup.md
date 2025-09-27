# CI/CD Setup Guide

This document explains how to set up CI/CD for the Blockchain-DApp project using GitHub Actions.

## Prerequisites

1. AWS Account (ID: 907849381252)
2. GitHub repository for the project
3. Appropriate IAM permissions for deployment

## GitHub Secrets Setup

To enable the CI/CD pipelines, you need to configure the following secrets in your GitHub repository:

### 1. AWS Credentials

Navigate to your GitHub repository settings > Secrets and variables > Actions, and add the following secrets:

- `AWS_ACCESS_KEY_ID` - Your AWS access key ID
- `AWS_SECRET_ACCESS_KEY` - Your AWS secret access key

These credentials should have sufficient permissions to:
- Create and manage S3 buckets
- Create and manage DynamoDB tables
- Create and manage EKS clusters
- Create and manage RDS instances
- Create and manage VPC resources

### 2. Application Deployment Secrets

For the application deployment workflow, you'll also need:

- `S3_BUCKET_NAME` - The S3 bucket name for hosting the frontend application
- `CLOUDFRONT_DISTRIBUTION_ID` - The CloudFront distribution ID for cache invalidation

### 3. Backend Deployment Secrets

For the backend deployment workflow, you'll need:

- `EKS_CLUSTER_NAME` - The name of the EKS cluster to deploy to

## IAM Policy for CI/CD

Create an IAM policy with the following permissions for your CI/CD pipeline:

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:*",
        "dynamodb:*",
        "eks:*",
        "rds:*",
        "ec2:*",
        "iam:*",
        "cloudformation:*",
        "cloudfront:CreateInvalidation"
      ],
      "Resource": "*"
    }
  ]
}
```

Note: In a production environment, you should restrict these permissions to only the specific resources needed.

## Workflow Triggers

The CI/CD workflows are configured to trigger on:

1. **Infrastructure Deployment** (`deploy-infra.yml`):
   - Push to `main` branch when files in the `infra/` directory are modified
   - Manual trigger via GitHub Actions UI

2. **Application Deployment** (`deploy-app.yml`):
   - Push to `main` branch when files in the `app/` directory are modified
   - Manual trigger via GitHub Actions UI

3. **Backend Deployment** (`deploy-backend.yml`):
   - Push to `main` branch when files in the `backend/` directory are modified
   - Manual trigger via GitHub Actions UI

## Deployment Process

1. Code is pushed to the repository
2. GitHub Actions workflow is triggered
3. Code is built and tested
4. Docker images are built and pushed to ECR
5. Kubernetes manifests are updated with new image tags
6. Changes are applied to the EKS cluster
7. CloudFront cache is invalidated (if needed)

## Required Permissions

The CI/CD pipeline requires the following AWS permissions:

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ecr:GetAuthorizationToken",
        "ecr:BatchCheckLayerAvailability",
        "ecr:GetDownloadUrlForLayer",
        "ecr:GetRepositoryPolicy",
        "ecr:DescribeRepositories",
        "ecr:ListImages",
        "ecr:DescribeImages",
        "ecr:BatchGetImage",
        "ecr:InitiateLayerUpload",
        "ecr:UploadLayerPart",
        "ecr:CompleteLayerUpload",
        "ecr:PutImage",
        "eks:DescribeCluster",
        "eks:ListClusters",
        "eks:AccessKubernetesApi",
        "eks:DescribeNodegroup",
        "eks:ListNodegroups",
        "eks:DescribeUpdate",
        "eks:ListUpdates",
        "eks:ListFargateProfiles",
        "eks:DescribeFargateProfile",
        "eks:ListIdentityProviderConfigs",
        "eks:DescribeIdentityProviderConfig",
        "eks:AssociateAccessPolicy",
        "s3:ListAllMyBuckets",
        "s3:ListBucket",
        "s3:GetBucketLocation",
        "s3:GetObject",
        "s3:PutObject",
        "s3:DeleteObject",
        "s3:CreateBucket",
        "s3:DeleteBucket",
        "cloudfront:CreateInvalidation",  // Added for CloudFront cache invalidation
        "cloudfront:GetInvalidation",
        "cloudfront:ListDistributions"
      ],
      "Resource": "*"
    }
  ]
}
```

## Deployment Order

The infrastructure deployment workflow ensures proper deployment order:

1. Global infrastructure (CloudTrail, Config, IAM, S3, DynamoDB)
2. Development environment
3. Staging environment (only on main branch)
4. Production environment (only on main branch with approval)

## Monitoring

The workflows include testing steps to ensure code quality before deployment:

- Unit tests for application code
- Terraform plan validation for infrastructure changes