#!/bin/bash

# Script to deploy Prometheus and Grafana monitoring stack

set -e

echo "Setting up Kubernetes context..."
aws eks update-kubeconfig --name blockchain-dev --region us-east-1

echo "Adding Helm repositories..."
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

echo "Creating monitoring namespace..."
kubectl create namespace monitoring || true

echo "Deploying Prometheus..."
helm upgrade --install prometheus prometheus-community/prometheus \
  --namespace monitoring \
  --values ../infra/env/dev/prometheus-values.yaml \
  --set server.serviceAccount.create=false \
  --set server.serviceAccount.name=prometheus-server \
  --set alertmanager.serviceAccount.create=false \
  --set alertmanager.serviceAccount.name=prometheus-server \
  --timeout 10m0s

echo "Deploying Grafana..."
helm upgrade --install grafana grafana/grafana \
  --namespace monitoring \
  --values ../infra/env/dev/grafana-values.yaml \
  --timeout 10m0s

echo "Waiting for pods to be ready..."
kubectl wait --for=condition=ready pod -l app.kubernetes.io/name=prometheus -n monitoring --timeout=300s
kubectl wait --for=condition=ready pod -l app.kubernetes.io/name=grafana -n monitoring --timeout=300s

echo "Getting service information..."
echo "Prometheus service:"
kubectl get svc prometheus-server -n monitoring

echo "Grafana service:"
kubectl get svc grafana -n monitoring

echo "Getting Grafana admin password..."
kubectl get secret --namespace monitoring grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

echo "Monitoring stack deployment completed!"