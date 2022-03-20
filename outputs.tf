output "service_url" {
  value = google_cloud_run_service.library.status[0].url
}
