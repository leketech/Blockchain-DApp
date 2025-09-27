output "endpoint" {
  description = "The connection endpoint"
  value       = aws_db_instance.main.endpoint
}

output "port" {
  description = "The database port"
  value       = aws_db_instance.main.port
}

output "db_name" {
  description = "The database name"
  value       = aws_db_instance.main.db_name
}

output "username" {
  description = "The master username"
  value       = aws_db_instance.main.username
  sensitive   = true
}

output "password" {
  description = "The master password"
  value       = aws_db_instance.main.password
  sensitive   = true
}

output "db_instance_identifier" {
  description = "The RDS instance identifier"
  value       = aws_db_instance.main.id
}

output "arn" {
  description = "The ARN of the RDS instance"
  value       = aws_db_instance.main.arn
}