provider "aws" {
  region = "us-east-1"
}

# Variables
variable "account_id" {
  description = "AWS Account ID"
  type        = string
  default     = "907849381252"
}

variable "acm_certificate_arn" {
  description = "ARN of the ACM certificate for HTTPS"
  type        = string
  default     = ""
}

variable "cloudfront_aliases" {
  description = "List of domain aliases for the CloudFront distribution"
  type        = list(string)
  default     = []
}

# VPC Module
module "vpc" {
  source = "../../modules/vpc"

  name = "blockchain-dev"

  cidr_block = "10.0.0.0/16"

  azs = ["us-east-1a", "us-east-1b", "us-east-1c"]

  public_subnets  = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  private_subnets = ["10.0.11.0/24", "10.0.12.0/24", "10.0.13.0/24"]
  database_subnets = ["10.0.21.0/24", "10.0.22.0/24", "10.0.23.0/24"]

  enable_nat_gateway   = true
  single_nat_gateway   = true # Cost optimization for dev
  enable_vpc_flow_logs = true

  tags = {
    Environment = "dev"
    Project     = "blockchain-dapp"
  }

  aws_region = "us-east-1"
}

# EKS Module
module "eks" {
  source = "../../modules/eks"

  name    = "blockchain-dev"
  vpc_id  = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  cluster_version = "1.27"

  endpoint_private_access = false
  endpoint_public_access  = true

  node_groups = {
    dev = {
      instance_types = ["t3.medium"]
      desired_size   = 2
      min_size       = 1
      max_size       = 3
      disk_size      = 20
      ami_type       = "AL2_x86_64"
    }
  }

  tags = {
    Environment = "dev"
    Project     = "blockchain-dapp"
  }
}

# RDS Module
module "rds" {
  source = "../../modules/rds-postgres"

  identifier              = "blockchain-dev"
  db_name                 = "blockchain_dev"
  username                = "blockchain"
  password                = "blockchain123" # In production, use AWS Secrets Manager
  db_subnet_group_name    = module.vpc.db_subnet_group_name
  vpc_security_group_ids  = [aws_security_group.rds.id]

  allocated_storage       = 20
  engine_version          = "13.7"
  instance_class          = "db.t3.micro" # Cost optimization for dev

  multi_az                = false # Cost optimization for dev
  backup_retention_period = 7
  skip_final_snapshot     = true
  deletion_protection     = false

  tags = {
    Environment = "dev"
    Project     = "blockchain-dapp"
  }
}

# Security Module
module "security" {
  source = "../../modules/security-module"

  vpc_id      = module.vpc.vpc_id
  environment = "dev"
  aws_region  = "us-east-1"
}

# Monitoring Module
module "monitoring" {
  source = "../../modules/monitoring-module"

  environment             = "dev"
  instance_id             = "i-1234567890abcdef0" # Placeholder, will be replaced with actual instance ID
  db_instance_identifier  = module.rds.identifier
  alarm_email             = "techlekedev@gmail.com"
  aws_region              = "us-east-1"
}

# Cost Monitoring Module
module "cost_monitoring" {
  source = "../../modules/cost-monitoring-module"

  environment                = "dev"
  monthly_budget_amount      = "1000"
  budget_notification_email  = "techlekedev@gmail.com"
  anomaly_threshold          = 100
  anomaly_notification_email = "techlekedev@gmail.com"
  reserve_instances          = false
  availability_zone          = "us-east-1a"
  instance_type              = "t3.medium"
  instance_count             = 2
  aws_region                 = "us-east-1"
  aws_account_id             = var.account_id
}

# PCI DSS Module
module "pci_dss" {
  source = "../../modules/pci-dss-module"

  environment              = "dev"
  config_role_arn          = "arn:aws:iam::907849381252:role/config-role" # Placeholder
  config_s3_bucket         = "config-bucket-dev-907849381252" # Placeholder
  cloudtrail_s3_bucket     = "cloudtrail-bucket-dev-907849381252" # Placeholder
  cloudtrail_kms_key_arn   = module.security.wallet_encryption_key_arn
  flow_log_role_arn        = "arn:aws:iam::907849381252:role/flow-log-role" # Placeholder
  flow_log_destination     = "arn:aws:logs:us-east-1:907849381252:log-group:/aws/flow-logs-dev" # Placeholder
  vpc_id                   = module.vpc.vpc_id
  aws_account_id           = var.account_id
  sensitive_data_bucket    = "sensitive-data-bucket-dev-907849381252" # Placeholder
}

# Security Group for RDS
resource "aws_security_group" "rds" {
  name        = "blockchain-dev-rds-sg"
  description = "Security group for RDS instance"
  vpc_id      = module.vpc.vpc_id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = [module.vpc.vpc_cidr_block]
    description = "PostgreSQL access from within VPC"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow all outbound traffic"
  }

  tags = {
    Name        = "blockchain-dev-rds-sg"
    Environment = "dev"
  }
}

# S3 Bucket for application data
resource "aws_s3_bucket" "app_data" {
  bucket = "dapp-bucket-2578"
}

resource "aws_s3_bucket_server_side_encryption_configuration" "app_data_encryption" {
  bucket = aws_s3_bucket.app_data.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_public_access_block" "app_data_public_block" {
  bucket = aws_s3_bucket.app_data.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# CloudFront Module
module "cloudfront" {
  source = "../../modules/cloudfront"

  environment             = "dev"
  s3_bucket_domain_name   = aws_s3_bucket.app_data.bucket_regional_domain_name
  acm_certificate_arn     = var.acm_certificate_arn
  aliases                 = var.cloudfront_aliases
}

# S3 Bucket Policy for CloudFront access
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

# KMS Key for application secrets
resource "aws_kms_key" "app_secrets" {
  description             = "KMS key for application secrets"
  deletion_window_in_days = 10
  enable_key_rotation     = true

  tags = {
    Name        = "app-secrets-kms-key"
    Environment = "dev"
  }
}

resource "aws_kms_alias" "app_secrets" {
  name          = "alias/app-secrets-key-dev"
  target_key_id = aws_kms_key.app_secrets.key_id
}

# IAM Role for EKS service account (IRSA)
resource "aws_iam_role" "eks_service_account" {
  name = "blockchain-dev-eks-service-account"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Federated = "arn:aws:iam::907849381252:oidc-provider/${replace(module.eks.cluster_oidc_issuer_url, "https://", "")}"
        }
        Action = "sts:AssumeRoleWithWebIdentity"
        Condition = {
          StringEquals = {
            "${replace(module.eks.cluster_oidc_issuer_url, "https://", "")}:sub" = "system:serviceaccount:monitoring:prometheus-server"
          }
        }
      }
    ]
  })
}

# IAM Policy for EKS service account
resource "aws_iam_policy" "eks_service_account" {
  name        = "blockchain-dev-eks-service-account-policy"
  description = "IAM policy for EKS service account"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "secretsmanager:GetSecretValue",
          "kms:Decrypt"
        ]
        Resource = [
          "*"
        ]
      }
    ]
  })
}

# Attach IAM policy to IAM role
resource "aws_iam_role_policy_attachment" "eks_service_account" {
  role       = aws_iam_role.eks_service_account.name
  policy_arn = aws_iam_policy.eks_service_account.arn
}

# Secrets Manager secret for database credentials
resource "aws_secretsmanager_secret" "db_credentials" {
  name = "blockchain-dev-db-credentials"
}

resource "aws_secretsmanager_secret_version" "db_credentials" {
  secret_id = aws_secretsmanager_secret.db_credentials.id
  secret_string = jsonencode({
    username = module.rds.username
    password = module.rds.password
    host     = module.rds.endpoint
    port     = module.rds.port
    dbname   = module.rds.db_name
  })
}

# Kubernetes provider configuration
provider "kubernetes" {
  host                   = module.eks.cluster_endpoint
  cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
  exec {
    api_version = "client.authentication.k8s.io/v1beta1"
    command     = "aws"
    args        = ["eks", "get-token", "--cluster-name", "blockchain-dev"]
  }
}

# Kubernetes namespace for monitoring
resource "kubernetes_namespace" "monitoring" {
  metadata {
    name = "monitoring"
  }
}

# Kubernetes namespace for the dApp
resource "kubernetes_namespace" "dapp" {
  metadata {
    name = "dapp"
  }
}

# Kubernetes service account for Prometheus
resource "kubernetes_service_account" "prometheus" {
  metadata {
    name      = "prometheus-server"
    namespace = kubernetes_namespace.monitoring.metadata[0].name
    annotations = {
      "eks.amazonaws.com/role-arn" = aws_iam_role.eks_service_account.arn
    }
  }
}

# Sample dApp Health Checker deployment
resource "kubernetes_deployment" "dapp_health_checker" {
  metadata {
    name      = "dapp-health-checker"
    namespace = kubernetes_namespace.dapp.metadata[0].name
    labels = {
      app = "dapp-health-checker"
    }
  }

  spec {
    replicas = 2

    selector {
      match_labels = {
        app = "dapp-health-checker"
      }
    }

    template {
      metadata {
        labels = {
          app = "dapp-health-checker"
        }
        annotations = {
          "prometheus.io/scrape" = "true"
          "prometheus.io/port"   = "8080"
          "prometheus.io/path"   = "/metrics"
        }
      }

      spec {
        container {
          image = "dapp-health-checker:latest"  # This will be replaced with the ECR image
          name  = "dapp-health-checker"

          port {
            container_port = 8080
          }

          env {
            name = "DB_HOST"
            value_from {
              secret_key_ref {
                name = "db-credentials"
                key  = "host"
              }
            }
          }

          env {
            name = "DB_PORT"
            value_from {
              secret_key_ref {
                name = "db-credentials"
                key  = "port"
              }
            }
          }

          env {
            name = "DB_USER"
            value_from {
              secret_key_ref {
                name = "db-credentials"
                key  = "username"
              }
            }
          }

          env {
            name = "DB_PASSWORD"
            value_from {
              secret_key_ref {
                name = "db-credentials"
                key  = "password"
              }
            }
          }

          env {
            name = "DB_NAME"
            value_from {
              secret_key_ref {
                name = "db-credentials"
                key  = "dbname"
              }
            }
          }

          env {
            name = "PORT"
            value = "8080"
          }

          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "0.25"
              memory = "256Mi"
            }
          }

          readiness_probe {
            http_get {
              path = "/health"
              port = 8080
            }
            initial_delay_seconds = 5
            period_seconds        = 10
          }

          liveness_probe {
            http_get {
              path = "/health"
              port = 8080
            }
            initial_delay_seconds = 30
            period_seconds        = 10
          }
        }
      }
    }
  }
}

# Kubernetes service for dApp Health Checker
resource "kubernetes_service" "dapp_health_checker" {
  metadata {
    name      = "dapp-health-checker"
    namespace = kubernetes_namespace.dapp.metadata[0].name
  }

  spec {
    selector = {
      app = "dapp-health-checker"
    }

    port {
      port        = 80
      target_port = 80
    }

    type = "ClusterIP"
  }
}

# Kubernetes secret for database credentials
resource "kubernetes_secret" "db_credentials" {
  metadata {
    name      = "db-credentials"
    namespace = kubernetes_namespace.dapp.metadata[0].name
  }

  data = {
    host     = module.rds.endpoint
    port     = module.rds.port
    username = module.rds.username
    password = module.rds.password
    dbname   = module.rds.db_name
  }
}

# Outputs
output "vpc_id" {
  description = "The ID of the VPC"
  value       = module.vpc.vpc_id
}

output "eks_cluster_endpoint" {
  description = "Endpoint for EKS control plane"
  value       = module.eks.cluster_endpoint
}

output "eks_cluster_certificate_authority_data" {
  description = "Base64 encoded certificate data required to communicate with the cluster"
  value       = module.eks.cluster_certificate_authority_data
}

output "rds_endpoint" {
  description = "The connection endpoint for the RDS instance"
  value       = module.rds.endpoint
}

output "app_data_bucket" {
  description = "The S3 bucket for application data"
  value       = aws_s3_bucket.app_data.bucket
}

output "wallet_encryption_key_arn" {
  description = "Wallet encryption key ARN"
  value       = module.security.wallet_encryption_key_arn
}

output "database_encryption_key_arn" {
  description = "Database encryption key ARN"
  value       = module.security.database_encryption_key_arn
}

output "logs_encryption_key_arn" {
  description = "Logs encryption key ARN"
  value       = module.security.logs_encryption_key_arn
}

# CloudFront outputs
output "cloudfront_domain_name" {
  description = "CloudFront distribution domain name"
  value       = module.cloudfront.cloudfront_domain_name
}

output "cloudfront_arn" {
  description = "CloudFront distribution ARN"
  value       = module.cloudfront.cloudfront_arn
}
