terraform {
  required_version = ">= 1.0.10"

  required_providers {
    nws = {
      source  = "nws/nws"
      version = "0.1.2"
    }
  }
}

module "instance" {
  source = "../../src"

  network_id     = var.network_id
  ip             = var.ip
  name           = var.name
  instance_type  = var.instance_type
  template       = var.template
  root_disk_size = var.root_disk_size
}
