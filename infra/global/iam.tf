# Create IAM groups for different roles
resource "aws_iam_group" "admins" {
  name = "admins"
}

resource "aws_iam_group" "developers" {
  name = "developers"
}

resource "aws_iam_group" "auditors" {
  name = "auditors"
}

# Create IAM policies
resource "aws_iam_policy" "admin_policy" {
  name        = "AdminPolicy"
  description = "Policy for administrators"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*"
    }
  ]
}
</EOF
}

resource "aws_iam_policy" "developer_policy" {
  name        = "DeveloperPolicy"
  description = "Policy for developers"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:*",
        "s3:*",
        "rds:*",
        "lambda:*",
        "logs:*",
        "iam:Get*",
        "iam:List*"
      ],
      "Resource": "*"
    }
  ]
}
</EOF
}

resource "aws_iam_policy" "auditor_policy" {
  name        = "AuditorPolicy"
  description = "Policy for auditors"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:Get*",
        "s3:List*",
        "config:Get*",
        "config:List*",
        "config:Describe*",
        "cloudtrail:Describe*",
        "cloudtrail:Get*",
        "cloudtrail:List*",
        "logs:Get*",
        "logs:Describe*",
        "logs:Filter*"
      ],
      "Resource": "*"
    }
  ]
}
</EOF
}

# Attach policies to groups
resource "aws_iam_group_policy_attachment" "admins_admin_policy" {
  group      = aws_iam_group.admins.name
  policy_arn = aws_iam_policy.admin_policy.arn
}

resource "aws_iam_group_policy_attachment" "developers_developer_policy" {
  group      = aws_iam_group.developers.name
  policy_arn = aws_iam_policy.developer_policy.arn
}

resource "aws_iam_group_policy_attachment" "auditors_auditor_policy" {
  group      = aws_iam_group.auditors.name
  policy_arn = aws_iam_policy.auditor_policy.arn
}

# Password policy for strong passwords
resource "aws_iam_account_password_policy" "strict" {
  minimum_password_length        = 12
  require_numbers                = true
  require_symbols                = true
  require_lowercase_characters   = true
  require_uppercase_characters   = true
  allow_users_to_change_password = true
  max_password_age               = 90
  password_reuse_prevention      = 5
}