variable "environment" {
  description = "Environment name"
  type        = string
}

variable "monthly_budget_amount" {
  description = "Monthly budget amount in USD"
  type        = string
}

variable "budget_notification_email" {
  description = "Email address for budget notifications"
  type        = string
}

variable "anomaly_threshold" {
  description = "Anomaly detection threshold"
  type        = number
}

variable "anomaly_notification_email" {
  description = "Email address for anomaly notifications"
  type        = string
}

variable "reserve_instances" {
  description = "Whether to reserve instances"
  type        = bool
}

variable "availability_zone" {
  description = "Availability zone for reserved instances"
  type        = string
}

variable "instance_type" {
  description = "Instance type for reserved instances"
  type        = string
}

variable "instance_count" {
  description = "Number of instances to reserve"
  type        = number
}

variable "aws_region" {
  description = "AWS region"
  type        = string
}

variable "aws_account_id" {
  description = "AWS account ID"
  type        = string
}