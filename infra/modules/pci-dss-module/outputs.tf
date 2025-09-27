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