
output "name" {
  description = "Instance names"
  value       = module.instance.name[*]
}

output "id" {
  description = "Instance IDs"
  value       = module.instance.id[*]
}
