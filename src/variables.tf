# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------
variable "network_id" {
  description = "Your subnet UUID, to which attach the instance"
  type        = string
}

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------
variable "zone" {
  description = "Your zone name"
  type        = string
  default     = "ru-msk-0"
}
variable "ip" {
  description = "Your VMs IP range"
  type        = list(string)
  default     = ["10.0.1.10"]
}
variable "name" {
  description = "Your instance names"
  type        = list(string)
  default     = ["VM0"]
}

variable "instance_type" {
  description = "Your instance types"
  type        = list(string)
  default     = ["t2.micro"]
}

variable "template" {
  description = "Your instance templates"
  type        = list(string)
  default     = ["Ubuntu 20.04 LTS"]
}

variable "root_disk_size" {
  description = "Your system disk size in GB"
  type        = list(number)
  default     = [10]
}

variable "keypair" {
  description = "Your SSH keypair to access the instance"
  type        = string
  default     = null
}

variable "tags" {
  description = "Your instance tags. Example: Type = Compute"
  type        = map(string)
  default     = null
}
variable "group" {
  description = "Your instance group name"
  type        = string
  default     = null
}
