terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 3.53"
    }
  }
}

provider "google" {
  project = var.project
}

locals {
  function_folder = "func-go"
  function_name   = "sorting-alpha"

  service_folder = "cloud-run-goapp"
  service_name   = "gabes-library"

  bucket_name   = "${var.project}"

  deployment_name = "library"
  gabes_library_worker_sa  = "serviceAccount:${google_service_account.gabes_library_worker.email}"
}