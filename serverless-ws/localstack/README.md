# [Localstack](https://www.localstack.cloud/)
Reason for using Localstack:
1. Localstack provides a fully functional local AWS cloud stack, which allows for testing AWS services without incurring costs or needing an internet connection.
2. It supports a wide range of AWS services, including S3, DynamoDB, Lambda, and more, making it versatile for various testing scenarios.


## [Installation](https://docs.localstack.cloud/aws/getting-started/installation/)
### Install Localstack CLI
To install the Localstack CLI, use home brew:
```bash
brew install localstack/localstack/localstack
```

There are several ways to install Localstack, including using Docker, home brew, pip, or the Localstack CLI. For my development environment, I use Docker or Docker Compose for ease of setup and isolation.
### Using Docker
1. Ensure you have Docker installed on your machine.
2. Pull the Localstack Docker image:
   ```bash
   docker pull localstack/localstack
   ```
3. Run Localstack using Docker:
   ```bash
   docker run -d -p 4566:4566 -p 4510-4559:4510-4559 localstack/localstack
   ```
4. Verify that Localstack is running by checking the logs:
   ```bash
   docker logs <container_id>
   ```
   
### docker-compose
Alternatively, you can use `docker-compose` to run Localstack. Create a `docker-compose.yml` file with the following content:
```yaml
services:
  localstack:
    container_name: "${LOCALSTACK_DOCKER_NAME:-localstack-main}"
    image: localstack/localstack
    ports:
      - "127.0.0.1:4566:4566"            # LocalStack Gateway
      - "127.0.0.1:4510-4559:4510-4559"  # external services port range
    environment:
      # LocalStack configuration: https://docs.localstack.cloud/references/configuration/
      - DEBUG=${DEBUG:-0}
    volumes:
      - "${LOCALSTACK_VOLUME_DIR:-./volume}:/var/lib/localstack"
      - "/var/run/docker.sock:/var/run/docker.sock"
```
Run Localstack with:
```bash
docker-compose up -d
```
This will start Localstack with the specified services and configurations.



