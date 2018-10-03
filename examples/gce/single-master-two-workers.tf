provider "stackpoint" {
  /* Set environment variable SPC_API_TOKEN with your API token from StackPointCloud    
     Set environment variable SPC_BASE_API_URL with API endpoint,   
     defaults to StackPointCloud production enviroment */
}

data "stackpoint_keysets" "keyset_default" {
  /* You can specify a custom orgID here,   
     or the system will find and use your default organization ID */
}

data "stackpoint_instance_specs" "master-specs" {
  provider_code = "${var.gce_code}"
  node_size     = "${var.gce_master_size}"
}

data "stackpoint_instance_specs" "worker-specs" {
  provider_code = "${var.gce_code}"
  node_size     = "${var.gce_worker_size}"
}

resource "stackpoint_cluster" "terraform-cluster" {
  org_id               = "${data.stackpoint_keysets.keyset_default.org_id}"
  cluster_name         = "Test GCE Cluster TerraForm2"
  provider_code        = "${var.gce_code}"
  provider_keyset      = "${data.stackpoint_keysets.keyset_default.gce_keyset}"
  region               = "${var.gce_region}"
  k8s_version          = "${var.gce_k8s_version}"
  startup_master_size  = "${data.stackpoint_instance_specs.master-specs.node_size}"
  startup_worker_count = 2
  startup_worker_size  = "${data.stackpoint_instance_specs.worker-specs.node_size}"
  region               = "${var.gce_region}"
  rbac_enabled         = true
  dashboard_enabled    = true
  etcd_type            = "classic"
  platform             = "${var.gce_platform}"
  channel              = "stable"
  ssh_keyset           = "${data.stackpoint_keysets.keyset_default.user_ssh_keyset}"
}