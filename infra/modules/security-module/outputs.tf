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