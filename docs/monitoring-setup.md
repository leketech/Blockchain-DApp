# Monitoring Setup Guide

This document explains how to set up monitoring for the Blockchain-DApp using Prometheus and Grafana.

## Overview

The monitoring stack consists of:
1. **Prometheus** - For metrics collection and alerting
2. **Grafana** - For visualization and dashboarding
3. **Alertmanager** - For handling alerts

## Components

### Prometheus
- Collects metrics from the dApp Health Checker and other Kubernetes components
- Stores metrics for 15 days
- Provides alerting rules for CPU and memory usage
- Integrated with Alertmanager for notification handling

### Grafana
- Provides dashboards for visualizing metrics
- Pre-configured with Kubernetes monitoring dashboards
- Connected to Prometheus as a data source
- Default admin credentials: admin/admin123

### Alertmanager
- Handles alerts from Prometheus
- Configured with Gmail SMTP for email notifications
- Supports webhook notifications for integration with other systems

## Deployment

### Prerequisites
1. EKS cluster must be running
2. kubectl configured to access the cluster
3. Helm 3 installed

### Deployment Steps
1. Run the deployment script:
   ```bash
   ./scripts/deploy-monitoring.sh
   ```

2. Or deploy manually using Helm:
   ```bash
   # Add Helm repositories
   helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
   helm repo add grafana https://grafana.github.io/helm-charts
   helm repo update
   
   # Deploy Prometheus
   helm install prometheus prometheus-community/prometheus \
     --namespace monitoring \
     --create-namespace \
     --values infra/env/dev/prometheus-values.yaml
   
   # Deploy Grafana
   helm install grafana grafana/grafana \
     --namespace monitoring \
     --values infra/env/dev/grafana-values.yaml
   ```

## Configuration

### Prometheus Configuration
The Prometheus configuration is defined in [prometheus-values.yaml](../infra/env/dev/prometheus-values.yaml):
- Scrapes metrics from the dApp Health Checker
- Includes alerting rules for CPU and memory usage
- Configured with Alertmanager for notifications

### Grafana Configuration
The Grafana configuration is defined in [grafana-values.yaml](../infra/env/dev/grafana-values.yaml):
- Pre-loaded with Kubernetes monitoring dashboards
- Connected to Prometheus as a data source
- Persistent storage for dashboards and settings

### Alertmanager Configuration
Alertmanager is configured in the Prometheus values file:
- Gmail SMTP configuration for email notifications
- Webhook support for integration with other systems
- Default routing to Gmail notifications

## Accessing the Services

### Prometheus
Port forward to access Prometheus:
```bash
kubectl port-forward -n monitoring svc/prometheus-server 9090:80
```
Then open http://localhost:9090

### Grafana
Port forward to access Grafana:
```bash
kubectl port-forward -n monitoring svc/grafana 3000:80
```
Then open http://localhost:3000
Default credentials: admin/admin123

### Alertmanager
Port forward to access Alertmanager:
```bash
kubectl port-forward -n monitoring svc/prometheus-alertmanager 9093:9093
```
Then open http://localhost:9093

## Alerts

The following alerts are configured:
1. **CPU Usage High** - Triggers when CPU usage exceeds 80%
2. **Memory Usage High** - Triggers when memory usage exceeds 80%
3. **Instance Down** - Triggers when an instance is unreachable

## Customization

### Adding New Dashboards
1. Create a new dashboard in Grafana UI
2. Export the dashboard as JSON
3. Add the JSON to the dashboards section in grafana-values.yaml

### Adding New Alerts
1. Add alerting rules to the rules section in prometheus-values.yaml
2. Update the alertmanager configuration if needed

### Changing Notification Settings
1. Update the Alertmanager configuration in prometheus-values.yaml
2. Change the SMTP settings or add webhook URLs