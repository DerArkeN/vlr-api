# VLR-API

## Description

The fastest VALORANT E-Sports API using data from [vlr.gg](https://vlr.gg). Supports gRPC and REST requests.

## Installation

To run this project, you need to have Docker installed on your machine. Follow either one of these steps:

### Pull from docker hub (recommended):

1. Pull the image: `docker pull derarken/vlr-api`
2. Run the image: `docker run -p 8080:8080 -p 8090:8090 derarken/vlr-api`

### Build yourself from the main branch:

1. Clone the repository: `git clone https://github.com/DerArkeN/vlr-api.git`
2. Navigate to the project directory: `cd vlr-api`
3. Build the Docker image: `docker build -t [image name] .`
4. Run the Docker container: `docker run -p 8080:8080 -p 8090:8090 [image name]`

## Usage

Once the Docker container is running, you can access the project using the following endpoints:

- gRPC: `localhost:8080`
- Swagger UI: `localhost:8090`
- REST: `localhost:8090/v1`

## Contributing

Contributions are welcome! If you would like to contribute to this project, please follow these steps:

1. Fork the repository
2. Create a new branch: `git checkout -b [branch name]`
3. Make your changes and commit them: `git commit -m '[commit message]'`
4. Push to the branch: `git push origin [branch name]`
5. Submit a pull request
