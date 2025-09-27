# S3 bucket for Terraform state files - Dev
resource "aws_s3_bucket" "terraform_state_dev" {
  bucket = "terraform-state-dev-907849381252"
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
  bucket = "terraform-state-staging-907849381252"
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
  bucket = "terraform-state-prod-907849381252"
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

# S3 bucket for application assets
resource "aws_s3_bucket" "app_assets" {
  bucket = "blockchain-dapp-assets-907849381252"
}

resource "aws_s3_bucket_versioning" "app_assets_versioning" {
  bucket = aws_s3_bucket.app_assets.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "app_assets_encryption" {
  bucket = aws_s3_bucket.app_assets.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "aws:kms"
      kms_master_key_id = aws_kms_key.s3.arn
    }
  }
}

resource "aws_s3_bucket_public_access_block" "app_assets_public_block" {
  bucket = aws_s3_bucket.app_assets.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# S3 bucket for logs
resource "aws_s3_bucket" "logs" {
  bucket = "blockchain-dapp-logs-907849381252"
}

resource "aws_s3_bucket_lifecycle_configuration" "logs_lifecycle" {
  bucket = aws_s3_bucket.logs.id

  rule {
    id     = "delete-after-365-days"
    status = "Enabled"

    expiration {
      days = 365
    }
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "logs_encryption" {
  bucket = aws_s3_bucket.logs.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "aws:kms"
      kms_master_key_id = aws_kms_key.s3.arn
    }
  }
}

resource "aws_s3_bucket_public_access_block" "logs_public_block" {
  bucket = aws_s3_bucket.logs.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}