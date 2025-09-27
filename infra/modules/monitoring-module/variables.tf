variable "environment" {
  description = "Environment name"
  type        = string
}

variable "instance_id" {
  description = "EC2 instance ID"
  type        = string
}

variable "db_instance_identifier" {
  description = "RDS instance identifier"
  type        = string
}

variable "alarm_email" {
  description = "Email address for alarm notifications"
  type        = string
}

variable "aws_region" {
  description = "AWS region"
  type        = string
}