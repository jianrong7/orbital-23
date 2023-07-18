# output "frontend_domain_name" {
#   description = "frontend domain name"
#   value       = "https://${aws_cloudfront_distribution.react_app_distribution.domain_name}"
# }

output "api_gw_public_ip" {
  description = "Public IPv4 address of API Gateway."
  value       = aws_instance.api_gateway.public_ip
}

output "api_gw_public_address" {
  description = "Public IPv4 address and port of API Gateway."
  value       = "${aws_instance.api_gateway.public_ip}:8888"
}

output "consul_server_public_ip" {
  description = "Public IPv4 address of Consul Server."
  value       = aws_instance.consul_server.public_ip
}

output "consul_server_public_address" {
  description = "Public IPv4 address and port of Consul Server."
  value       = "${aws_instance.consul_server.public_ip}:8500"
}

output "consul_server_private_ip" {
  description = "Private IPv4 address of Consul Server."
  value       = aws_instance.consul_server.private_ip
}

output "consul_server_private_address" {
  description = "Private IPv4 address of Consul Server."
  value       = "${aws_instance.consul_server.private_ip}:8500"
}

