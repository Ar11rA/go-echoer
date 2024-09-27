# Go Echo Server

## Project Integrations

| **Service**         | **Details**                                           |
|---------------------|-------------------------------------------------------|
| **File Management**  | - Upload and save files                               |
|                     | - Read files from a directory                         |
|                     | - File validation and error handling                  |
| **Redis**           | - Connect to Redis                                    |
|                     | - Perform CRUD operations in Redis                    |
|                     | - Cache data using Redis                              |
|                     | - Handle Redis connection lifecycle                   |
| **PostgreSQL (PSQL)**| - Establish a PostgreSQL connection                   |
|                     | - Perform database migrations                         |
|                     | - CRUD operations with PostgreSQL                     |
|                     | - Query optimization and indexing                     |
|                     | - Handle database connection pooling                  |
| **HTTP Server**      | - Handle HTTP requests and responses                  |
|                     | - Define RESTful routes for different services        |
|                     | - Log requests and responses in JSON format           |
|                     | - Middleware for request validation and security      |
| **MongoDB**          | - Connect to MongoDB                                  |
|                     | - Perform CRUD operations in MongoDB                  |
|                     | - Handle MongoDB connection pooling                   |
|                     | - Schema design and indexing                          |
| **RabbitMQ**         | - Connect to RabbitMQ                                 |
|                     | - Publish/Subscribe messaging                         |
|                     | - Message queue management                            |
|                     | - Handle RabbitMQ connection lifecycle                |
| **MinIO**            | - Connect to MinIO                                    |
|                     | - Upload and retrieve objects from MinIO              |
|                     | - Handle file storage in MinIO buckets                |
|                     | - Manage MinIO connection lifecycle                   |

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
4. Test the server
   ```bash
   go test -v ./...
   ```
5. Run the server
   ```bash
   go run .
   ```

## Contributing

Contributions are welcome! Please create a pull request or open an issue for any improvements or bug reports.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

You can copy this directly into a `README.md` file! Let me know if you need further modifications.
