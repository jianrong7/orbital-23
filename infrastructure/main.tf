terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

resource "random_uuid" "uuid" {}

provider "aws" {
  region  = "ap-southeast-1"
  # profile = "default"
  access_key = AWS_ACCESS_KEY
  secret_key = AWS_SECRET_KEY
}