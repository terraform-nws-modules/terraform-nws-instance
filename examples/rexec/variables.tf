# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------
variable "network_id" {
  type = string
}

variable "ip" {
  type = list(string)
}

variable "name" {
  type = list(string)
}

variable "instance_type" {
  type = list(string)
}

variable "template" {
  type = list(string)
}

variable "root_disk_size" {
  type = list(number)
}

# variable "public_ip" {
#   type    = string
#   default = "10.0.2.20"
# }

variable "ssh_user" {
  type    = string
  default = "admin"
}

variable "ssh_port" {
  type    = number
  default = 3659
}
