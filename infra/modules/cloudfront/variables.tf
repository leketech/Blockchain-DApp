variable "environment" {
  description = "Environment name (dev, staging, prod)"
  type        = string
}

variable "s3_bucket_domain_name" {
  description = "S3 bucket domain name for CloudFront origin"
  type        = string
}

variable "acm_certificate_arn" {
  description = "ARN of the ACM certificate for HTTPS"
  type        = string
  default     = ""
}

variable "aliases" {
  description = "List of domain aliases for the CloudFront distribution"
  type        = list(string)
  default     = []
}