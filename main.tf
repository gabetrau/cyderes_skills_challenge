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
  service_folder = "goservice"
  service_name   = "library"

  bucket_name   = "${var.project}-cyderes"

  deployment_name = "library"
  library_worker_sa = "serviceAccount:${google_service_account.library_worker.email}"
}
