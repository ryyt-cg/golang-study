# Terraform Workspace
This directory contains Terraform configurations for managing cloud resources.

## Prerequisites
- [Terraform](https://www.terraform.io/downloads.html) installed on your machine.
- [AWS CLI](https://aws.amazon.com/cli/) configured with your AWS credentials.
- (Optional) [Localstack](https://localstack.cloud/) for local development and testing.
- (Optional) [Docker](https://www.docker.com/) for running Localstack or other services.
- (Optional) [Terraform Cloud](https://www.terraform.io/cloud) for remote state management and collaboration.

## Project List
1. All Terraform projects in this directory are organized by their respective cloud services. Each project contains its own `main.tf`, `variables.tf`, and `outputs.tf` files, along with any necessary modules or resources.
2. Use tflocal to manage the Terraform state and configurations.
3. Use tflocal plan, apply, and destroy commands to manage the resources to localstack.



