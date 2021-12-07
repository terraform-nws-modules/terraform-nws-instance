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
  zone             = var.zone
  name             = var.name[count.index]
  service_offering = var.instance_type[count.index]
  network_id       = var.network_id
  ip_address       = var.ip[count.index]
  template         = var.template[count.index]
  root_disk_size   = var.root_disk_size[count.index]
  group            = var.group
  keypair          = var.keypair
  expunge          = true
  start_vm         = true
  tags             = var.tags

  provisioner "remote-exec" {
    inline = [
      "sudo hostnamectl set-hostname my-vm",
      "echo Done!"
    ]
    connection {
      type     = "ssh"
      port     = 3659
      user     = "admin"
      password = "admin"
      host     = ""
      timeout  = "45s"
    }
  }
}
