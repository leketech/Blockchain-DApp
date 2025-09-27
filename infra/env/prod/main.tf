provider "aws" {
  region = "us-east-1"
}

# Variables
variable "account_id" {
  description = "AWS Account ID"
  type        = string
  default     = "907849381252"
}