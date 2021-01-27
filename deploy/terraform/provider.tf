provider "google" {
  project     = "pipedrive-nagy"
  region      = "us-east1"
  zone        = "us-east1-b"
  credentials = file("../credentials/serviceaccount.json")
}