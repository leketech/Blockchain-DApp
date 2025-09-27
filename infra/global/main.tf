provider "aws" {
  region = var.aws_region
}

# Variables
variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "account_id" {
  description = "AWS Account ID"
  type        = string
}

# S3 bucket for Terraform state files - Dev
resource "aws_s3_bucket" "terraform_state_dev" {
  bucket = "terraform-state-dev-${var.account_id}"
}

resource "aws_s3_bucket_versioning" "terraform_state_dev_versioning" {
  bucket = aws_s3_bucket.terraform_state_dev.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "terraform_state_dev_encryption" {
  bucket = aws_s3_bucket.terraform_state_dev.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_public_access_block" "terraform_state_dev_public_block" {
  bucket = aws_s3_bucket.terraform_state_dev.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# S3 bucket for Terraform state files - Staging
resource "aws_s3_bucket" "terraform_state_staging" {
  bucket = "terraform-state-staging-${var.account_id}"
}

resource "aws_s3_bucket_versioning" "terraform_state_staging_versioning" {
  bucket = aws_s3_bucket.terraform_state_staging.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "terraform_state_staging_encryption" {
  bucket = aws_s3_bucket.terraform_state_staging.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_public_access_block" "terraform_state_staging_public_block" {
  bucket = aws_s3_bucket.terraform_state_staging.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# S3 bucket for Terraform state files - Prod
resource "aws_s3_bucket" "terraform_state_prod" {
  bucket = "terraform-state-prod-${var.account_id}"
}

resource "aws_s3_bucket_versioning" "terraform_state_prod_versioning" {
  bucket = aws_s3_bucket.terraform_state_prod.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "terraform_state_prod_encryption" {
  bucket = aws_s3_bucket.terraform_state_prod.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_public_access_block" "terraform_state_prod_public_block" {
  bucket = aws_s3_bucket.terraform_state_prod.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# DynamoDB table for Terraform state locking - Dev
resource "aws_dynamodb_table" "terraform_state_lock_dev" {
  name         = "terraform-state-lock-dev"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Name = "terraform-state-lock-dev"
  }
}

# DynamoDB table for Terraform state locking - Staging
resource "aws_dynamodb_table" "terraform_state_lock_staging" {
  name         = "terraform-state-lock-staging"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Name = "terraform-state-lock-staging"
  }
}

# DynamoDB table for Terraform state locking - Prod
resource "aws_dynamodb_table" "terraform_state_lock_prod" {
  name         = "terraform-state-lock-prod"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Name = "terraform-state-lock-prod"
  }
}