# CloudFront Setup for Blockchain DApp

This document explains how to set up and configure CloudFront CDN for the Blockchain DApp.

## Overview

CloudFront is a content delivery network (CDN) service that securely delivers data, videos, applications, and APIs to customers globally with low latency and high transfer speeds. In this project, we use CloudFront to serve static assets from the S3 bucket.

The distribution will be identified as "Dapp-2578" in the comments and tags for easy identification.

## Module Structure

The CloudFront module is located at `infra/modules/cloudfront` and contains:

- `main.tf`: Main CloudFront distribution configuration
- `variables.tf`: Input variables for the module
- `outputs.tf`: Output values from the module

## Configuration

### Variables

The CloudFront module accepts the following variables:

- `environment`: Environment name (dev, staging, prod)
- `s3_bucket_domain_name`: S3 bucket domain name for CloudFront origin
- `acm_certificate_arn`: (Optional) ARN of the ACM certificate for HTTPS
- `aliases`: (Optional) List of domain aliases for the CloudFront distribution

### Usage in Environment

To use the CloudFront module in an environment (e.g., dev), add the following to your `main.tf`:

```hcl
module "cloudfront" {
  source = "../../modules/cloudfront"

  environment             = "dev"
  s3_bucket_domain_name   = aws_s3_bucket.app_data.bucket_regional_domain_name
  acm_certificate_arn     = var.acm_certificate_arn  # Optional
  aliases                 = var.cloudfront_aliases  # Optional
}
```

The S3 bucket used for CloudFront origin is named `dapp-bucket-2578`.

### S3 Bucket Policy

The module requires an S3 bucket policy to allow CloudFront access:

```hcl
resource "aws_s3_bucket_policy" "app_data_cloudfront" {
  bucket = aws_s3_bucket.app_data.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "AllowCloudFrontServicePrincipalReadOnly"
        Effect    = "Allow"
        Principal = {
          AWS = module.cloudfront.origin_access_identity
        }
        Action    = [
          "s3:GetObject"
        ]
        Resource  = "arn:aws:s3:::${aws_s3_bucket.app_data.bucket}/*"
      }
    ]
  })
}
```

## Deployment

1. Navigate to the environment directory:
   ```bash
   cd infra/env/dev
   ```

2. Initialize Terraform:
   ```bash
   terraform init
   ```

3. Plan the deployment:
   ```bash
   terraform plan
   ```

4. Apply the changes:
   ```bash
   terraform apply
   ```

## HTTPS Configuration

For production environments, you should configure HTTPS using an ACM certificate:

1. Request an ACM certificate in the AWS console
2. Add the certificate ARN to your Terraform variables:
   ```bash
   terraform apply -var="acm_certificate_arn=arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012"
   ```

## Custom Domain Configuration

To use a custom domain with CloudFront:

1. Add the domain to your Terraform variables:
   ```bash
   terraform apply -var='cloudfront_aliases=["app.example.com"]'
   ```

2. Create a CNAME record in your DNS provider pointing to the CloudFront domain name

## Outputs

The module provides the following outputs:

- `cloudfront_domain_name`: The CloudFront distribution domain name
- `cloudfront_hosted_zone_id`: The CloudFront distribution hosted zone ID
- `cloudfront_arn`: The CloudFront distribution ARN
- `origin_access_identity`: The CloudFront origin access identity path

## Useful Scripts

The project includes helpful scripts for managing the Dapp-2578 CloudFront distribution:

1. **invalidate-cloudfront.sh**: Invalidates the CloudFront cache
   ```bash
   ./scripts/invalidate-cloudfront.sh <distribution-id> [paths]
   ```

2. **find-dapp-2578-distribution.sh**: Helps identify the Dapp-2578 distribution
   ```bash
   ./scripts/find-dapp-2578-distribution.sh
   ```

## Troubleshooting

### Common Issues

1. **Access Denied Errors**: Ensure the S3 bucket policy is correctly configured to allow CloudFront access.

2. **Certificate Issues**: Make sure the ACM certificate is in the us-east-1 region.

3. **Domain Validation**: For custom domains, ensure DNS records are correctly configured.

### Useful Commands

Check CloudFront distribution status:
```bash
aws cloudfront get-distribution --id <distribution-id>
```

Create an invalidation to refresh cached content:
```bash
aws cloudfront create-invalidation --distribution-id <distribution-id> --paths "/*"