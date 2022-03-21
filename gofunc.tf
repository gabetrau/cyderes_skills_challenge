# The Cloud Function
resource "google_cloudfunctions_function" "function" {
  name        = local.gofunc_name
  description = "sort titles from json array"
  runtime     = "go113"
  region      = var.region

  available_memory_mb   = 128
  trigger_http          = true
  entry_point           = "sortAlpha"
  service_account_email = google_service_account.library_worker.email

  depends_on = [google_project_service.cloudfunctions]
}