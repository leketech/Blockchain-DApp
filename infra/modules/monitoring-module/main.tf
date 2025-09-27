# Monitoring Module

# CloudWatch Alarms
resource "aws_cloudwatch_metric_alarm" "high_cpu_utilization" {
  alarm_name          = "high-cpu-utilization-${var.environment}"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/EC2"
  period              = "120"
  statistic           = "Average"
  threshold           = "80"
  alarm_description   = "This metric monitors EC2 instance CPU utilization"
  alarm_actions       = [aws_sns_topic.alarms.arn]

  dimensions = {
    InstanceId = var.instance_id
  }
}

resource "aws_cloudwatch_metric_alarm" "high_memory_utilization" {
  alarm_name          = "high-memory-utilization-${var.environment}"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "MemoryUtilization"
  namespace           = "AWS/EC2"
  period              = "120"
  statistic           = "Average"
  threshold           = "80"
  alarm_description   = "This metric monitors EC2 instance memory utilization"
  alarm_actions       = [aws_sns_topic.alarms.arn]

  dimensions = {
    InstanceId = var.instance_id
  }
}

resource "aws_cloudwatch_metric_alarm" "high_disk_utilization" {
  alarm_name          = "high-disk-utilization-${var.environment}"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "DiskUtilization"
  namespace           = "AWS/EC2"
  period              = "120"
  statistic           = "Average"
  threshold           = "80"
  alarm_description   = "This metric monitors EC2 instance disk utilization"
  alarm_actions       = [aws_sns_topic.alarms.arn]

  dimensions = {
    InstanceId = var.instance_id
  }
}

resource "aws_cloudwatch_metric_alarm" "database_cpu_utilization" {
  alarm_name          = "database-cpu-utilization-${var.environment}"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/RDS"
  period              = "120"
  statistic           = "Average"
  threshold           = "75"
  alarm_description   = "This metric monitors RDS instance CPU utilization"
  alarm_actions       = [aws_sns_topic.alarms.arn]

  dimensions = {
    DBInstanceIdentifier = var.db_instance_identifier
  }
}

resource "aws_cloudwatch_metric_alarm" "database_free_storage" {
  alarm_name          = "database-free-storage-${var.environment}"
  comparison_operator = "LessThanThreshold"
  evaluation_periods  = "1"
  metric_name         = "FreeStorageSpace"
  namespace           = "AWS/RDS"
  period              = "300"
  statistic           = "Average"
  threshold           = "5368709120" # 5GB in bytes
  alarm_description   = "This metric monitors RDS instance free storage space"
  alarm_actions       = [aws_sns_topic.alarms.arn]

  dimensions = {
    DBInstanceIdentifier = var.db_instance_identifier
  }
}

# SNS Topics for alerts
resource "aws_sns_topic" "alarms" {
  name = "alarms-${var.environment}"

  tags = {
    Name = "alarms-topic-${var.environment}"
  }
}

resource "aws_sns_topic_subscription" "email_subscription" {
  topic_arn = aws_sns_topic.alarms.arn
  protocol  = "email"
  endpoint  = var.alarm_email
}

# CloudWatch Dashboards
resource "aws_cloudwatch_dashboard" "main_dashboard" {
  dashboard_name = "main-dashboard-${var.environment}"

  dashboard_body = jsonencode({
    widgets = [
      {
        type   = "metric"
        x      = 0
        y      = 0
        width  = 12
        height = 6

        properties = {
          metrics = [
            ["AWS/EC2", "CPUUtilization", "InstanceId", var.instance_id]
          ]
          period = 300
          stat   = "Average"
          region = var.aws_region
          title  = "EC2 CPU Utilization"
        }
      },
      {
        type   = "metric"
        x      = 12
        y      = 0
        width  = 12
        height = 6

        properties = {
          metrics = [
            ["AWS/RDS", "CPUUtilization", "DBInstanceIdentifier", var.db_instance_identifier]
          ]
          period = 300
          stat   = "Average"
          region = var.aws_region
          title  = "RDS CPU Utilization"
        }
      }
    ]
  })
}

# Outputs
output "alarms_topic_arn" {
  value = aws_sns_topic.alarms.arn
}

output "dashboard_url" {
  value = "https://console.aws.amazon.com/cloudwatch/home?region=${var.aws_region}#dashboards:name=${aws_cloudwatch_dashboard.main_dashboard.dashboard_name}"
}