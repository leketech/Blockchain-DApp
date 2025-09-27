# KMS Key for S3 encryption
resource "aws_kms_key" "s3" {
  description             = "KMS key for S3 bucket encryption"
  deletion_window_in_days = 10
  enable_key_rotation     = true

  tags = {
    Name        = "s3-kms-key"
    Environment = "global"
  }
}

resource "aws_kms_alias" "s3" {
  name          = "alias/s3-key"
  target_key_id = aws_kms_key.s3.key_id
}

# KMS Key for EKS secrets
resource "aws_kms_key" "eks_secrets" {
  description             = "KMS key for EKS secrets encryption"
  deletion_window_in_days = 10
  enable_key_rotation     = true

  tags = {
    Name        = "eks-secrets-kms-key"
    Environment = "global"
  }
}

resource "aws_kms_alias" "eks_secrets" {
  name          = "alias/eks-secrets-key"
  target_key_id = aws_kms_key.eks_secrets.key_id
}