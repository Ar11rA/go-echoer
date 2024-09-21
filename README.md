# Go Echo Server

This project is a simple RESTful API server built using the Echo framework in Go. It includes various features such as health checks, file operations, and HTTP integrations. The server is structured to follow best practices, including dependency injection and organized routes.

## Project Structure

```sh
.
├── Dockerfile
├── Dockerfile.qa
├── README.md
├── config
│   └── logger.go
├── docker-compose-e2e-test.yaml
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
├── routes
│   ├── file.go
│   ├── health.go
│   ├── http.go
│   └── routes.go
├── server.go
├── services
│   ├── container.go
│   ├── file_service.go
│   ├── file_service_test.go
│   ├── http_service.go
│   └── http_service_test.go
├── types
│   └── response.go
└── utils
    ├── file_util.go
    └── http_util.go
```

## Getting Started

### Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go**: Install Go version 1.22 or later. You can download it from the official [Go website](https://golang.org/dl/).
- **Docker**: Install Docker to containerize the application. Download it from the [Docker website](https://www.docker.com/products/docker-desktop).
- **Docker Compose**: Install Docker Compose to manage multi-container applications. You can find installation instructions [here](https://docs.docker.com/compose/install/).
- **PostgreSQL**: If your application uses a PostgreSQL database, ensure it is set up and running. You can use Docker to run a PostgreSQL container for testing.


### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Ar11rA/quote-server.git
   ```

2. Enter into quote-server: `cd quote-server`
3. Install Dependencies
   ```bash
   go mod tidy
   go install
   ```
4. Run the server
   ```bash
   go run .
   ```

## Contributing

Contributions are welcome! Please create a pull request or open an issue for any improvements or bug reports.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

You can copy this directly into a `README.md` file! Let me know if you need further modifications.
