variable "environment" {
  description = "Environment name"
  type        = string
}

variable "config_role_arn" {
  description = "Config role ARN"
  type        = string
}

variable "config_s3_bucket" {
  description = "Config S3 bucket name"
  type        = string
}

variable "cloudtrail_s3_bucket" {
  description = "CloudTrail S3 bucket name"
  type        = string
}

variable "cloudtrail_kms_key_arn" {
  description = "CloudTrail KMS key ARN"
  type        = string
}

variable "flow_log_role_arn" {
  description = "Flow log role ARN"
  type        = string
}

variable "flow_log_destination" {
  description = "Flow log destination"
  type        = string
}

variable "vpc_id" {
  description = "VPC ID"
  type        = string
}

variable "aws_account_id" {
  description = "AWS account ID"
  type        = string
}

variable "sensitive_data_bucket" {
  description = "Sensitive data bucket name"
  type        = string
}