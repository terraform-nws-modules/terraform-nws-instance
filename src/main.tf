terraform {
  required_version = ">= 1.0.10"

  required_providers {
    nws = {
      source  = "nws/nws"
      version = "0.1.2"
    }
  }
}

resource "nws_instance" "inst" {
  count            = length(var.name)
  group            = var.group
  name             = var.name[count.index]
  service_offering = var.instance_type[count.index]
  zone             = var.zone
  template         = var.template[count.index]
  network_id       = var.network_id
  ip_address       = var.ip[count.index]
  root_disk_size   = var.root_disk_size[count.index]
  expunge          = true
}
