# Blockchain-DApp

This is a decentralized application with a complete infrastructure setup on AWS.

## Key Technology Choices

The project uses a modern technology stack carefully selected for performance, scalability, and developer productivity:

- **Frontend**: React + React Native for unified web and mobile development
- **Backend**: Golang for high-performance blockchain operations
- **Database**: PostgreSQL for ACID-compliant ledger entries
- **Infrastructure**: EKS (Kubernetes) for container orchestration
- **Infrastructure as Code**: Terraform for reproducible environments
- **Card Services**: Stripe Issuing / Marqeta for card issuing and tokenization
- **Observability**: Prometheus + Grafana + Alertmanager stack
- **CI/CD**: GitHub Actions for automated workflows

See [Technology Choices & Rationale](docs/architecture/technology-choices.md) for detailed explanations of each technology decision.

## AWS Account Setup

The AWS account (ID: 907849381252) has been configured with:

1. **CloudTrail** - For logging all AWS API calls across all regions
2. **Config** - For tracking configuration changes of AWS resources
3. **IAM Security** - With groups, policies, and strong password policies
4. **Secure Root Account** - With proper separation of duties

## VPC Topology & Network Plan

The infrastructure follows a well-designed VPC topology with:

- **Public Subnets** - For internet-facing resources
- **Private Subnets** - For application resources
- **Database Subnets** - For RDS instances
- **NAT Gateways** - For private subnet internet access
- **VPC Endpoints** - For private access to AWS services
- **VPC Flow Logs** - For network traffic monitoring

See [VPC Design Document](docs/vpc-design.md) for detailed network architecture.

## Infrastructure Components

### Terraform Modules
- **VPC Module** - Creates a complete VPC with public/private subnets, NAT gateways, and routing
- **EKS Module** - Creates EKS clusters with node groups
- **RDS Module** - Creates PostgreSQL database instances

### Environments
- `dev` - Development environment
- `staging` - Staging environment
- `prod` - Production environment

### Global Resources
- CloudTrail for API logging
- Config for resource tracking
- IAM policies and groups
- S3 buckets for Terraform state and application assets
- DynamoDB tables for state locking
- KMS keys for encryption
- ECR repositories for container images

## Frontend Applications

### Web Interface
The dApp includes a responsive web interface built with React and TypeScript:

- Modern UI with dark/light mode support
- Wallet creation and import functionality
- Responsive design for all screen sizes
- Material Symbols icons
- Tailwind CSS styling

See [Web Interface Documentation](web/README.md) for build and deployment instructions.

### Mobile Applications
The dApp also includes mobile applications for iOS and Android:

- Native mobile experience
- Wallet creation and import functionality
- Platform-specific UI components

See [Mobile App Documentation](mobile/README.md) for build and deployment instructions.

## CI/CD Setup

The project uses GitHub Actions for continuous integration and deployment:

- **Infrastructure Deployment** - Automatically deploys infrastructure changes
- **Application Deployment** - Builds and deploys the frontend application
- **Backend Deployment** - Builds and deploys the backend services

See [CI/CD Setup Guide](docs/ci-cd-setup.md) for detailed configuration instructions.

## Monitoring and Observability

The dev environment includes a complete monitoring stack:

- **Prometheus** - For metrics collection and alerting
- **Grafana** - For visualization and dashboarding
- **Alertmanager** - For handling alerts with Gmail SMTP or webhook notifications
- **dApp Health Checker** - Sample application with Prometheus metrics

See [Monitoring Setup Guide](docs/monitoring-setup.md) for detailed configuration instructions.

## Deployment

Use the Makefile commands to deploy:

```bash
# Deploy global resources
make deploy-global

# Deploy environment-specific resources
make deploy-dev
make deploy-staging
make deploy-prod
```

Or use GitHub Actions for automated deployment.

## Running the Project Locally

You can run the entire project locally using Docker Compose:

1. Make sure you have Docker and Docker Compose installed
2. Run the project:
   ```bash
   ./run-project.sh
   ```
   
   Or manually with Docker Compose:
   ```bash
   # For newer versions of Docker (v20.10+)
   docker compose up --build
   
   # For older versions of Docker
   docker-compose up --build
   ```

3. Access the services:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - Database: postgres://blockchain:blockchain@localhost:5432/blockchain_dev

4. To stop the services:
   ```bash
   # For newer versions of Docker (v20.10+)
   docker compose down
   
   # For older versions of Docker
   docker-compose down
   ```

## Running in Development Mode

For development, you can run the frontend and backend separately:

### Prerequisites
- Go 1.19+
- Node.js 16+
- PostgreSQL

### Setting up the Database
```bash
./setup-db.sh
```

### Running the Backend
```bash
./run-backend-dev.sh
```

### Running the Frontend
```bash
./run-frontend-dev.sh
```

## Security Features

- CloudTrail logs are stored in a secure S3 bucket with lifecycle policies (365 days)
- Config records all resource configurations with delivery to S3
- IAM groups with least-privilege policies (admins, developers, auditors)
- Strong password policy enforced (12+ characters, numbers, symbols, uppercase, lowercase)
- Separate Terraform state storage per environment with versioning and encryption
- Public access blocked on all S3 buckets
- VPC Flow Logs for network monitoring
- DynamoDB locking for Terraform state consistency
- KMS encryption for data at rest
- IRSA (IAM Roles for Service Accounts) for secure Kubernetes access
- Secrets Manager for sensitive data