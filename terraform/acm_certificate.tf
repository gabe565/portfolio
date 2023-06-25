resource "aws_acm_certificate" "app" {
  provider = aws.us-east-1

  domain_name       = "gabecook.com"
  validation_method = "DNS"

  subject_alternative_names = ["*.gabecook.com"]

  tags = {
    app = "portfolio"
  }

  lifecycle {
    create_before_destroy = true
  }
}

output "cert_validation_records" {
  value       = aws_acm_certificate.app.domain_validation_options
  description = "ACM certificate validation DNS records"
}
