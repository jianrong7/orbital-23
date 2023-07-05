#########################
# CloudFront distribution
#########################
resource "aws_cloudfront_distribution" "react_app_distribution" {
  enabled             = true
  is_ipv6_enabled     = true
  comment             = "React App Distribution"
  default_root_object = "index.html"

  default_cache_behavior {
    target_origin_id = "S3Origin"
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD"]
    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }
    viewer_protocol_policy = "redirect-to-https"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }

  origin {
    origin_id   = "S3Origin"
    domain_name = aws_s3_bucket.frontend_bucket.bucket_regional_domain_name
    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "http-only"
      origin_ssl_protocols   = ["TLSv1.2"]
    }
  }

  tags = {
    Name = "React App"
  }
}