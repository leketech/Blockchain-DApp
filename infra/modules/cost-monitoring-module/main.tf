# Cost Monitoring Module

# Budget for the environment
resource "aws_budgets_budget" "monthly_budget" {
  name              = "monthly-budget-${var.environment}"
  budget_type       = "COST"
  limit_amount      = var.monthly_budget_amount
  limit_unit        = "USD"
  time_period_start = "2023-01-01_00:00"
  time_unit         = "MONTHLY"

  notification {
    comparison_operator        = "GREATER_THAN"
    threshold                  = 80
    threshold_type             = "PERCENTAGE"
    notification_type          = "ACTUAL"
    subscriber_email_addresses = [var.budget_notification_email]
  }

  notification {
    comparison_operator        = "GREATER_THAN"
    threshold                  = 100
    threshold_type             = "PERCENTAGE"
    notification_type          = "ACTUAL"
    subscriber_email_addresses = [var.budget_notification_email]
  }

  cost_filter {
    name   = "Service"
    values = ["AmazonEC2", "AmazonRDS", "AmazonS3", "AmazonKMS", "AWSCloudTrail"]
  }
}

# Cost Anomaly Detection
resource "aws_ce_anomaly_monitor" "main_monitor" {
  name              = "main-anomaly-monitor-${var.environment}"
  monitor_type      = "DIMENSIONAL"
  monitor_dimension = "SERVICE"
}

resource "aws_ce_anomaly_subscription" "main_subscription" {
  name      = "main-anomaly-subscription-${var.environment}"
  threshold = var.anomaly_threshold
  frequency = "IMMEDIATE"

  monitor_arn_list = [
    aws_ce_anomaly_monitor.main_monitor.arn
  ]

  subscriber {
    address = var.anomaly_notification_email
    type    = "EMAIL"
  }
}

# Reserved Instance Recommendations
resource "aws_ec2_capacity_reservation" "reserved_instances" {
  count = var.reserve_instances ? 1 : 0

  availability_zone = var.availability_zone
  instance_type     = var.instance_type
  instance_platform = "Linux/UNIX"
  instance_count    = var.instance_count
  tenancy           = "default"
  ebs_optimized     = true

  tags = {
    Name = "reserved-instances-${var.environment}"
  }
}

# Outputs
output "budget_name" {
  value = aws_budgets_budget.monthly_budget.name
}

output "anomaly_monitor_arn" {
  value = aws_ce_anomaly_monitor.main_monitor.arn
}