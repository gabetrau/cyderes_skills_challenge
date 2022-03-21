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
  function_folder = "gofunc"
  function_name = "sortAlpha"

  service_folder = "goservice"
  service_name   = "library"

  deployment_name = "library"
  library_worker_sa = "serviceAccount:${google_service_account.library_worker.email}"
}
