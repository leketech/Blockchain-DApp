terraform {
  backend "s3" {
    bucket         = "terraform-state-dev-907849381252"
    key            = "dev/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform-state-lock-dev"
    encrypt        = true
  }
}