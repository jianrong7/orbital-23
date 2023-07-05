output "frontend_domain_name" {
  description = "frontend domain name"
  value       = aws_cloudfront_distribution.react_app_distribution.domain_name
}