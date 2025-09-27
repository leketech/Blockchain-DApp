# Makefile for deploying AWS infrastructure

# Variables
ACCOUNT_ID := 907849381252
AWS_REGION := us-east-1

# Deploy infrastructure for a specific environment
deploy-dev:
	cd infra/env/dev && terraform init && terraform apply -auto-approve

deploy-staging:
	cd infra/env/staging && terraform init && terraform apply -auto-approve

deploy-prod:
	cd infra/env/prod && terraform init && terraform apply -auto-approve

# Deploy global infrastructure
deploy-global:
	cd infra/global && terraform init && terraform apply -auto-approve

# Plan infrastructure changes for a specific environment
plan-dev:
	cd infra/env/dev && terraform plan

plan-staging:
	cd infra/env/staging && terraform plan

plan-prod:
	cd infra/env/prod && terraform plan

# Plan global infrastructure changes
plan-global:
	cd infra/global && terraform plan

# Destroy infrastructure for a specific environment (use with caution)
destroy-dev:
	cd infra/env/dev && terraform destroy -auto-approve

destroy-staging:
	cd infra/env/staging && terraform destroy -auto-approve

destroy-prod:
	cd infra/env/prod && terraform destroy -auto-approve

# Destroy global infrastructure (use with caution)
destroy-global:
	cd infra/global && terraform destroy -auto-approve

# Build and push the health checker image
build-health-checker:
	./scripts/build-health-checker.sh

# Deploy monitoring stack
deploy-monitoring:
	./scripts/deploy-monitoring.sh

# Build web interface
build-web:
	./scripts/build-web.sh

# Test web interface
test-web:
	./scripts/test-web.sh

# Initialize mobile project
init-mobile:
	./scripts/init-mobile.sh

# Test mobile project
test-mobile:
	./scripts/test-mobile.sh

# Port forward to access services
port-forward-prometheus:
	kubectl port-forward -n monitoring svc/prometheus-server 9090:80

port-forward-grafana:
	kubectl port-forward -n monitoring svc/grafana 3000:80

port-forward-alertmanager:
	kubectl port-forward -n monitoring svc/prometheus-alertmanager 9093:9093

.PHONY: deploy-dev deploy-staging deploy-prod deploy-global \
        plan-dev plan-staging plan-prod plan-global \
        destroy-dev destroy-staging destroy-prod destroy-global \
        build-health-checker deploy-monitoring \
        build-web test-web \
        init-mobile test-mobile \
        port-forward-prometheus port-forward-grafana port-forward-alertmanager