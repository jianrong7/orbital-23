output "api_gw_public_address" {
  description = "Public IPv4 address and port of API Gateway."
  value       = "${aws_instance.api_gateway.public_ip}:8888"
}

output "consul_server_public_address" {
  description = "Public IPv4 address and port of Consul Server."
  value       = "${aws_instance.consul_server.public_ip}:8500"
}

output "api_gw_public_ip" {
  description = "Public IPv4 address of API Gateway."
  value       = aws_instance.api_gateway.public_ip
}

output "consul_server_public_ip" {
  description = "Public IPv4 address of Consul Server."
  value       = aws_instance.consul_server.public_ip
}

output "consul_server_private_address" {
  description = "Private IPv4 address of Consul Server."
  value       = "${aws_instance.consul_server.private_ip}:8500"
}

output "idl_management_service_public_ip" {
  description = "Public IPv4 address of IDL Management Service."
  value       = aws_instance.idl_management.public_ip
}

output "idl_management_service_private_address" {
  description = "Private IPv4 address of IDL Management Service."
  value       = "${aws_instance.idl_management.private_ip}:9999"
}

output "service1v1_public_ip" {
  description = "List of private IP addresses assigned to the service1v1 instances."
  value       = [aws_instance.service1v1.*.public_ip]
}

output "service1v2_public_ip" {
  description = "List of private IP addresses assigned to the service1v2 instances."
  value       = [aws_instance.service1v2.*.public_ip]
}

output "service2v1_public_ip" {
  description = "List of private IP addresses assigned to the service2v1 instances."
  value       = [aws_instance.service2v1.*.public_ip]
}

