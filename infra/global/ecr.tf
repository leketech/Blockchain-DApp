# ECR repository for blockchain backend
resource "aws_ecr_repository" "blockchain_backend" {
  name                 = "blockchain-backend"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  encryption_configuration {
    encryption_type = "KMS"
    kms_key = aws_kms_key.eks_secrets.arn
  }

  tags = {
    Environment = "dev"
    Service     = "blockchain-backend"
  }
}

# ECR repository for dApp Health Checker
resource "aws_ecr_repository" "dapp_health_checker" {
  name                 = "dapp-health-checker"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  encryption_configuration {
    encryption_type = "KMS"
    kms_key = aws_kms_key.eks_secrets.arn
  }

  tags = {
    Environment = "dev"
    Service     = "dapp-health-checker"
  }
}