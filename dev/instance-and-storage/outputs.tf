output "public_ip" {
  description = "name of created instances. "
  value       =  module.instance_and_storage.public_ip
}