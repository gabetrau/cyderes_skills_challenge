# The Cloud Run service
resource "google_cloud_run_service" "library" {
  name                       = local.service_name
  location                   = var.region
  autogenerate_revision_name = true

  template {
    spec {
      containers {
	image = "gcr.io/${var.project}/${local.service_name}" 
      }
      timeout_seconds = 900
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }

  depends_on = [google_project_service.run]
}

# Set service public
data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.library.location
  project  = google_cloud_run_service.library.project
  service  = google_cloud_run_service.library.name

  policy_data = data.google_iam_policy.noauth.policy_data
  depends_on  = [google_cloud_run_service.library]
}


