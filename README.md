# IPResist

## Overview

IPResist is a Go-based application that performs periodic health checks on IP addresses and stores the results in InfluxDB. The application uses MySQL for persistent data storage and Docker Compose for orchestration.

## Features

- Periodic health checks on specified IP addresses
- Storage of health check results in InfluxDB
- MySQL database for persistent data storage
- Configuration through environment variables

## Requirements

- Docker
- Docker-Compose

## Setup

### 1. Create a `.env` File

After clone the repositry, create a `.env` file in the root directory with the following content:

```env
ADDRESS=0.0.0.0
PORT=8080
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=mydb
DB_HOST=mysql
DB_PORT=3306
INFLUXDB_HOST=http://influxdb:8086
INFLUXDB_TOKEN=adminpassword
INFLUXDB_ORG=myorg
INFLUXDB_BUCKET=mybucket
HEALTH_CHECK_INTERVAL=300
```

### 2. Build and Start the Containers

Use Docker Compose to build and start the services:

```sh
docker-compose build .
docker-compose up -d
```

This will start the following services:

- MySQL (port 3306)
- InfluxDB (port 8086)
- IPResist Go application (port 8080)

### 3. Stop the Containers

To stop the running containers:

```sh
docker-compose down
```

## Configuration

Configuration for the application is done through environment variables. These can be set in the `.env` file in the root of the project. The following variables are used:

- `ADDRESS`: The address the Go application will listen on (default `0.0.0.0`)
- `PORT`: The port the Go application will listen on (default `8080`)
- `DB_USER`: MySQL database user
- `DB_PASSWORD`: MySQL database password
- `DB_NAME`: MySQL database name
- `DB_HOST`: MySQL database host (default `mysql`)
- `DB_PORT`: MySQL database port (default `3306`)
- `INFLUXDB_HOST`: InfluxDB host URL
- `INFLUXDB_TOKEN`: InfluxDB token
- `INFLUXDB_ORG`: InfluxDB organization
- `INFLUXDB_BUCKET`: InfluxDB bucket
- `HEALTH_CHECK_INTERVAL`: Interval for performing health checks (in seconds, default `300`)

## Running Health Checks

The application performs periodic health checks on the specified IP addresses using the scheduler defined in `internal/scheduler/scheduler.go`. The results are stored in InfluxDB.

## Contribution

Contributions are welcome! If you would like to contribute.

### Frontend Developer Needed

> I am looking for a front-end developer, preferably experienced with React, to create a web interface for our API. This interface will allow users to interact with the IPResist application and visualize health check results.

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details.

## Maintainer

Mohammad Reza Fadaei - [mohrezfadaei@gmail.com](mailto:mohrezfadaei@gmail.com)
