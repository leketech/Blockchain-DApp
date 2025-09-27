resource "aws_config_configuration_recorder" "main" {
  name     = "main-config-recorder"
  role_arn = aws_iam_role.config_role.arn

  recording_group {
    all_supported                 = true
    include_global_resource_types = true
  }
}

resource "aws_config_delivery_channel" "main" {
  name           = "main-config-delivery-channel"
  s3_bucket_name = aws_s3_bucket.config_bucket.bucket
  sns_topic_arn  = aws_sns_topic.config_topic.arn

  snapshot_delivery_properties {
    delivery_frequency = "Six_Hours"
  }

  depends_on = [
    aws_config_configuration_recorder.main
  ]
}

resource "aws_s3_bucket" "config_bucket" {
  bucket = "config-logs-${var.account_id}"
  force_destroy = true
}

resource "aws_s3_bucket_lifecycle_configuration" "config_bucket_lifecycle" {
  bucket = aws_s3_bucket.config_bucket.id

  rule {
    id     = "delete-after-365-days"
    status = "Enabled"

    expiration {
      days = 365
    }
  }
}

resource "aws_sns_topic" "config_topic" {
  name = "config-topic"
}

resource "aws_iam_role" "config_role" {
  name = "config-role"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "config.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
</POLICY>
}

resource "aws_iam_role_policy_attachment" "config_role_policy" {
  role       = aws_iam_role.config_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWS_ConfigRole"
}