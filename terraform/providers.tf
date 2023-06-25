terraform {
  backend "s3" {
    bucket = "gabe565-terraform"
    key    = "portfolio.tfstate"
    region = "us-east-2"
  }

  required_providers {
    aws = {
      version = "5.5.0"
    }
  }
}

provider "aws" {
  alias  = "us-east-1"
  region = "us-east-1"
}

provider "aws" {
  region = "us-east-2"
}

data "aws_caller_identity" "current" {}