output "cloudfront_domain_name" {
  description = "CloudFront distribution domain name"
  value       = aws_cloudfront_distribution.app_assets.domain_name
}

output "cloudfront_hosted_zone_id" {
  description = "CloudFront distribution hosted zone ID"
  value       = aws_cloudfront_distribution.app_assets.hosted_zone_id
}

output "cloudfront_arn" {
  description = "CloudFront distribution ARN"
  value       = aws_cloudfront_distribution.app_assets.arn
}

output "origin_access_identity" {
  description = "CloudFront origin access identity"
  value       = aws_cloudfront_origin_access_identity.app_assets.cloudfront_access_identity_path
}