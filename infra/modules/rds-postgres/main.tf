# RDS PostgreSQL Module

# Variables
variable "identifier" {
  description = "The name of the RDS instance"
  type        = string
}

variable "db_name" {
  description = "The name of the database to create when the DB instance is created"
  type        = string
}

variable "username" {
  description = "Username for the master DB user"
  type        = string
}

variable "password" {
  description = "Password for the master DB user"
  type        = string
  sensitive   = true
}

variable "db_subnet_group_name" {
  description = "Name of DB subnet group"
  type        = string
}

variable "vpc_security_group_ids" {
  description = "VPC security group IDs"
  type        = list(string)
}

variable "allocated_storage" {
  description = "The allocated storage in gigabytes"
  type        = number
  default     = 20
}

variable "engine" {
  description = "The database engine to use"
  type        = string
  default     = "postgres"
}

variable "engine_version" {
  description = "The engine version to use"
  type        = string
  default     = "13.7"
}

variable "instance_class" {
  description = "The instance type of the RDS instance"
  type        = string
  default     = "db.t3.micro"
}

variable "parameter_group_name" {
  description = "Name of the DB parameter group to associate"
  type        = string
  default     = "default.postgres13"
}

variable "availability_zone" {
  description = "The availability zone for the RDS instance"
  type        = string
  default     = null
}

variable "multi_az" {
  description = "Specifies if the RDS instance is multi-AZ"
  type        = bool
  default     = false
}

variable "backup_retention_period" {
  description = "The days to retain backups for"
  type        = number
  default     = 7
}

variable "backup_window" {
  description = "The daily time range during which automated backups are created"
  type        = string
  default     = "03:00-04:00"
}

variable "maintenance_window" {
  description = "The window to perform maintenance in"
  type        = string
  default     = "sun:04:00-sun:05:00"
}

variable "auto_minor_version_upgrade" {
  description = "Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window"
  type        = bool
  default     = true
}

variable "skip_final_snapshot" {
  description = "Determines whether a final DB snapshot is created before the DB instance is deleted"
  type        = bool
  default     = true
}

variable "deletion_protection" {
  description = "If the DB instance should have deletion protection enabled"
  type        = bool
  default     = false
}

variable "tags" {
  description = "Tags to apply to resources"
  type        = map(string)
  default     = {}
}

# RDS Instance
resource "aws_db_instance" "main" {
  identifier                = var.identifier
  db_name                   = var.db_name
  username                  = var.username
  password                  = var.password
  db_subnet_group_name      = var.db_subnet_group_name
  vpc_security_group_ids    = var.vpc_security_group_ids
  allocated_storage         = var.allocated_storage
  engine                    = var.engine
  engine_version            = var.engine_version
  instance_class            = var.instance_class
  parameter_group_name      = var.parameter_group_name
  availability_zone         = var.availability_zone
  multi_az                  = var.multi_az
  backup_retention_period   = var.backup_retention_period
  backup_window             = var.backup_window
  maintenance_window        = var.maintenance_window
  auto_minor_version_upgrade = var.auto_minor_version_upgrade
  skip_final_snapshot       = var.skip_final_snapshot
  deletion_protection       = var.deletion_protection

  tags = merge(
    var.tags,
    {
      Name = var.identifier
    }
  )
}

# Outputs
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