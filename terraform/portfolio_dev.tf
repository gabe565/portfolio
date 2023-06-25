module "dev" {
  source = "./modules/portfolio"

  aws_account_id      = data.aws_caller_identity.current.account_id
  env                 = "dev"
  hostname            = "dev.gabecook.com"
  acm_certificate_arn = aws_acm_certificate.app.arn
}

output "dev_cloudfront_hostname" {
  value       = module.dev.distribution_hostname
  description = "Prod Cloudfront hostname"
}
