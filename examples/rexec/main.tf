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

# ---------------------------------------------------------------------------------------------------------------------
# Provision the server using remote-exec
# ---------------------------------------------------------------------------------------------------------------------

resource "null_resource" "example_provisioner" {
  triggers = {
    public_ip = var.ip[0]
  }

  connection {
    type  = "ssh"
    host  = var.ip[0]
    user  = var.ssh_user
    port  = var.ssh_port
    agent = true
  }

  // copy our example script to the server
  provisioner "file" {
    source      = "files/getpath.sh"
    destination = "/tmp/getpath.sh"
  }

  // change permissions to executable and pipe its output into a new file
  provisioner "remote-exec" {
    inline = [
      "chmod +x /tmp/getpath.sh",
      "/tmp/getpath > /tmp/curr-path",
    ]
  }

  provisioner "local-exec" {
    # copy the public-ip file back to CWD, which will be tested
    command = "scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${var.ssh_user}@${var.ip[0]}:/tmp/curr-path curr-path"
  }
}
