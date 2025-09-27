output "budget_name" {
  value = aws_budgets_budget.monthly_budget.name
}

output "anomaly_monitor_arn" {
  value = aws_ce_anomaly_monitor.main_monitor.arn
}