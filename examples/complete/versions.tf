terraform {
  required_version = ">= 1.9.6"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 6.5.0"
    }
  }
}

provider "google" {
  region  = local.region
  project = local.project
}
