terraform {
  backend "s3" {
    bucket         = "terraform-state-staging-907849381252"
    key            = "staging/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform-state-lock-staging"
    encrypt        = true
  }
}