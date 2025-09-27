# Security Module

# KMS Keys for various services
resource "aws_kms_key" "wallet_encryption" {
  description             = "KMS key for wallet encryption"
  deletion_window_in_days = 30
  enable_key_rotation     = true

  tags = {
    Name = "wallet-encryption-key"
  }
}

resource "aws_kms_key" "database_encryption" {
  description             = "KMS key for database encryption"
  deletion_window_in_days = 30
  enable_key_rotation     = true

  tags = {
    Name = "database-encryption-key"
  }
}

resource "aws_kms_key" "logs_encryption" {
  description             = "KMS key for logs encryption"
  deletion_window_in_days = 30
  enable_key_rotation     = true

  tags = {
    Name = "logs-encryption-key"
  }
}

resource "aws_kms_alias" "wallet_encryption_alias" {
  name          = "alias/wallet-encryption"
  target_key_id = aws_kms_key.wallet_encryption.key_id
}

resource "aws_kms_alias" "database_encryption_alias" {
  name          = "alias/database-encryption"
  target_key_id = aws_kms_key.database_encryption.key_id
}

resource "aws_kms_alias" "logs_encryption_alias" {
  name          = "alias/logs-encryption"
  target_key_id = aws_kms_key.logs_encryption.key_id
}

# Security Groups
resource "aws_security_group" "app_sg" {
  name        = "app-security-group"
  description = "Security group for application instances"
  vpc_id      = var.vpc_id

  ingress {
    description = "HTTPS"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "app-security-group"
  }
}

resource "aws_security_group" "database_sg" {
  name        = "database-security-group"
  description = "Security group for database instances"
  vpc_id      = var.vpc_id

  ingress {
    description     = "PostgreSQL"
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.app_sg.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "database-security-group"
  }
}

resource "aws_security_group" "redis_sg" {
  name        = "redis-security-group"
  description = "Security group for Redis instances"
  vpc_id      = var.vpc_id

  ingress {
    description     = "Redis"
    from_port       = 6379
    to_port         = 6379
    protocol        = "tcp"
    security_groups = [aws_security_group.app_sg.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "redis-security-group"
  }
}

# Network ACLs
resource "aws_network_acl" "main_acl" {
  vpc_id = var.vpc_id

  egress {
    protocol   = "tcp"
    rule_no    = 100
    action     = "allow"
    cidr_block = "0.0.0.0/0"
    from_port  = 443
    to_port    = 443
  }

  ingress {
    protocol   = "tcp"
    rule_no    = 100
    action     = "allow"
    cidr_block = "0.0.0.0/0"
    from_port  = 443
    to_port    = 443
  }

  tags = {
    Name = "main-network-acl"
  }
}

# WAF Web ACL
resource "aws_wafv2_web_acl" "main_waf" {
  name        = "main-web-acl"
  description = "Main WAF Web ACL"
  scope       = "REGIONAL"

  default_action {
    allow {}
  }

  visibility_config {
    cloudwatch_metrics_enabled = true
    metric_name                = "main-web-acl"
    sampled_requests_enabled   = true
  }

  tags = {
    Name = "main-web-acl"
  }
}

# CloudWatch Log Groups with encryption
resource "aws_cloudwatch_log_group" "app_logs" {
  name              = "/aws/app/logs"
  retention_in_days = 30
  kms_key_id        = aws_kms_key.logs_encryption.arn

  tags = {
    Name = "app-logs"
  }
}

resource "aws_cloudwatch_log_group" "audit_logs" {
  name              = "/aws/audit/logs"
  retention_in_days = 90
  kms_key_id        = aws_kms_key.logs_encryption.arn

  tags = {
    Name = "audit-logs"
  }
}

# Outputs
output "wallet_encryption_key_arn" {
  value = aws_kms_key.wallet_encryption.arn
}

output "database_encryption_key_arn" {
  value = aws_kms_key.database_encryption.arn
}

output "logs_encryption_key_arn" {
  value = aws_kms_key.logs_encryption.arn
}

output "app_security_group_id" {
  value = aws_security_group.app_sg.id
}

output "database_security_group_id" {
  value = aws_security_group.database_sg.id
}

output "redis_security_group_id" {
  value = aws_security_group.redis_sg.id
}

output "waf_web_acl_arn" {
  value = aws_wafv2_web_acl.main_waf.arn
}

output "app_log_group_name" {
  value = aws_cloudwatch_log_group.app_logs.name
}

output "audit_log_group_name" {
  value = aws_cloudwatch_log_group.audit_logs.name
}