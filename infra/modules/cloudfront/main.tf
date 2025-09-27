# CloudFront distribution for serving static assets
resource "aws_cloudfront_distribution" "app_assets" {
  origin {
    domain_name = var.s3_bucket_domain_name
    origin_id   = "S3-${var.s3_bucket_domain_name}"

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.app_assets.cloudfront_access_identity_path
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  comment             = "Dapp-2578 - CloudFront distribution for ${var.environment} environment"
  default_root_object = "index.html"

  aliases = var.aliases

  default_cache_behavior {
    allowed_methods        = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = "S3-${var.s3_bucket_domain_name}"
    compress               = true
    viewer_protocol_policy = "redirect-to-https"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    min_ttl     = 0
    default_ttl = 3600
    max_ttl     = 86400
  }

  ordered_cache_behavior {
    path_pattern           = "/api/*"
    allowed_methods        = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = "S3-${var.s3_bucket_domain_name}"
    compress               = true
    viewer_protocol_policy = "https-only"

    forwarded_values {
      query_string = true

      headers = [
        "Authorization",
        "CloudFront-Viewer-Country",
        "CloudFront-Forwarded-Proto",
        "CloudFront-Is-Desktop-Viewer",
        "CloudFront-Is-Mobile-Viewer",
        "CloudFront-Is-Tablet-Viewer",
      ]

      cookies {
        forward = "all"
      }
    }

    min_ttl     = 0
    default_ttl = 0
    max_ttl     = 0
  }

  price_class = "PriceClass_100"

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  tags = {
    Environment = var.environment
    Project     = "blockchain-dapp"
    Name        = "Dapp-2578-cloudfront-${var.environment}"
  }

  dynamic "viewer_certificate" {
    for_each = var.acm_certificate_arn != "" ? [1] : []

    content {
      acm_certificate_arn      = var.acm_certificate_arn
      ssl_support_method       = "sni-only"
      minimum_protocol_version = "TLSv1.2_2021"
    }
  }

  dynamic "viewer_certificate" {
    for_each = var.acm_certificate_arn == "" ? [1] : []

    content {
      cloudfront_default_certificate = true
    }
  }
}

# CloudFront Origin Access Identity
resource "aws_cloudfront_origin_access_identity" "app_assets" {
  comment = "Dapp-2578 - OAI for blockchain-dapp-${var.environment} S3 bucket"
}