output "service_url" {
  value = google_cloud_run_service.gabes-library.status[0].url
}