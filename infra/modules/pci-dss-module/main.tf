# PCI DSS Compliance Module

# Enable AWS Config for compliance monitoring
resource "aws_config_configuration_recorder" "main" {
  name     = "pci-dss-config-recorder-${var.environment}"
  role_arn = var.config_role_arn

  recording_group {
    all_supported                 = true
    include_global_resource_types = true
  }
}

resource "aws_config_delivery_channel" "main" {
  name           = "pci-dss-delivery-channel-${var.environment}"
  s3_bucket_name = var.config_s3_bucket

  snapshot_delivery_properties {
    delivery_frequency = "TwentyFour_Hours"
  }

  depends_on = [aws_config_configuration_recorder.main]
}

# Enable CloudTrail for audit logging
resource "aws_cloudtrail" "pci_trail" {
  name                          = "pci-dss-trail-${var.environment}"
  s3_bucket_name                = var.cloudtrail_s3_bucket
  include_global_service_events = true
  is_multi_region_trail         = true
  enable_log_file_validation    = true
  kms_key_id                    = var.cloudtrail_kms_key_arn

  event_selector {
    read_write_type           = "All"
    include_management_events = true

    data_resource {
      type   = "AWS::S3::Object"
      values = ["arn:aws:s3:::${var.sensitive_data_bucket}/"]
    }
  }

  tags = {
    Name = "pci-dss-trail-${var.environment}"
  }
}

# Enable VPC Flow Logs for network monitoring
resource "aws_flow_log" "pci_flow_logs" {
  iam_role_arn    = var.flow_log_role_arn
  log_destination = var.flow_log_destination
  traffic_type    = "ALL"
  vpc_id          = var.vpc_id

  tags = {
    Name = "pci-dss-flow-logs-${var.environment}"
  }
}

# Enable GuardDuty for threat detection
resource "aws_guardduty_detector" "pci_detector" {
  enable = true

  datasources {
    s3_logs {
      enable = true
    }
  }

  tags = {
    Name = "pci-dss-detector-${var.environment}"
  }
}

# Enable Security Hub for security best practices
resource "aws_securityhub_account" "pci_securityhub" {
}

resource "aws_securityhub_standards_subscription" "pci_standard" {
  standards_arn = "arn:aws:securityhub:::ruleset/pci-dss/v/3.2.1"

  depends_on = [aws_securityhub_account.pci_securityhub]
}

# Enable Macie for sensitive data discovery
resource "aws_macie2_account" "pci_macie" {
  status = "ENABLED"
}

resource "aws_macie2_classification_job" "sensitive_data_job" {
  job_type = "ONE_TIME"
  name     = "pci-dss-sensitive-data-job-${var.environment}"

  s3_job_definition {
    bucket_definitions {
      account_id = var.aws_account_id
      buckets    = [var.sensitive_data_bucket]
    }
  }

  depends_on = [aws_macie2_account.pci_macie]
}

# Enable Inspector for vulnerability assessment
resource "aws_inspector2_enabler" "pci_inspector" {
  account_ids = [var.aws_account_id]

  for_each = toset(["EC2", "ECR", "LAMBDA"])
  resource_types = [each.key]
}

# Outputs
output "config_recorder_name" {
  value = aws_config_configuration_recorder.main.name
}

output "cloudtrail_name" {
  value = aws_cloudtrail.pci_trail.name
}

output "guardduty_detector_id" {
  value = aws_guardduty_detector.pci_detector.id
}

output "securityhub_account_id" {
  value = aws_securityhub_account.pci_securityhub.id
}