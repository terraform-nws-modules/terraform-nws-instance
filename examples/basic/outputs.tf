
output "name" {
  description = "Instance names"
  value       = instance.name[*]
}

output "id" {
  description = "Instance IDs"
  value       = instance.id[*]
}
