module "prod" {
  source = "./modules/portfolio"

  aws_account_id      = data.aws_caller_identity.current.account_id
  env                 = "prod"
  hostname            = "gabecook.com"
  acm_certificate_arn = aws_acm_certificate.app.arn
}

output "prod_cloudfront_hostname" {
  value       = module.prod.distribution_hostname
  description = "Prod Cloudfront hostname"
}
