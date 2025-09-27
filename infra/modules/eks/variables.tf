variable "name" {
  description = "Name to be used for EKS cluster and resource naming"
  type        = string
}

variable "vpc_id" {
  description = "VPC ID where EKS cluster will be created"
  type        = string
}

variable "subnet_ids" {
  description = "Subnet IDs for EKS cluster"
  type        = list(string)
}

variable "cluster_version" {
  description = "Kubernetes version for EKS cluster"
  type        = string
  default     = "1.27"
}

variable "endpoint_private_access" {
  description = "Enable private access to EKS cluster endpoint"
  type        = bool
  default     = false
}

variable "endpoint_public_access" {
  description = "Enable public access to EKS cluster endpoint"
  type        = bool
  default     = true
}

variable "node_groups" {
  description = "Map of node group configurations"
  type = map(object({
    instance_types = list(string)
    desired_size   = number
    min_size       = number
    max_size       = number
    disk_size      = number
    ami_type       = string
  }))
  default = {
    default = {
      instance_types = ["t3.medium"]
      desired_size   = 2
      min_size       = 1
      max_size       = 3
      disk_size      = 20
      ami_type       = "AL2_x86_64"
    }
  }
}

variable "tags" {
  description = "Tags to apply to resources"
  type        = map(string)
  default     = {}
}