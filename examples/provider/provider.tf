terraform {
  required_providers {
    segment_public_api = {
      source  = "ds-terraform/segment_public_api"
      version = "0.8.1"
    }
  }
}

provider "segment_public_api" {
  # Configuration options
}