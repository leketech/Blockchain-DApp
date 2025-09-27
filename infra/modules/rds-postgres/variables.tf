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