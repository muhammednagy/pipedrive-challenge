resource "google_container_cluster" "gke-cluster" {
  name     = "pipedrive-cluster"
  location = "us-east1"
  remove_default_node_pool = true
  initial_node_count = 1
  node_config {
    machine_type = "e2-micro"
  }
  timeouts {
    create = "30m"
    update = "20m"
  }
}
// using seperate node pool because it's "recommended" from terraform documentation.
resource "google_container_node_pool" "my-pipedrive-node-pool" {
  name       = "my-pipedrive-node-pool"
  location   = "us-east1"
  cluster    = google_container_cluster.gke-cluster.name
  initial_node_count = 2
  node_config {
    machine_type = "e2-small"
  }
  timeouts {
    create = "30m"
    update = "20m"
  }
}
