# FileStore

FileStore is a distributed file storage system built with Golang. It provides functionality for uploading, storing, and downloading files by splitting them into chunks and handling them in parallel. The system is designed to be highly scalable and fault-tolerant.

## Features

- **Upload Files**: Upload files by splitting them into chunks and saving them in parallel.
- **Download Files**: Retrieve and merge file chunks to reconstruct the original file.
- **File Metadata Management**: Store and retrieve file metadata.
- **Parallel Processing**: Efficiently handle file chunks using goroutines.

## Project Structure

- `api/models/`: Contains data models used in the application.
- `services/`: Contains business logic and service layer functions.
- `utils/`: Contains utility functions, including file processing and ID generation.
- `config/`: Configuration settings and database initialization.
- `main.go`: Entry point for the application.
- `Dockerfile`: Dockerfile for containerizing the application.
- `docker-compose.yml`: Docker Compose configuration for setting up the application stack.
- `tests/`: Contains unit, integration, and end-to-end tests.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.19 or higher)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup

### Local Development

1. **Clone the Repository**

```bash
git clone https://github.com/yourusername/filestore.git
cd filestore
```

2. **Install Dependencies**

Run the following command to download and install the required Go modules:

```bash
go mod download
```

3. **Run the Application**

You can start the application by running:

```bash

go run main.go
```

### Using Docker

1 **Build and Run with Docker Compose**

Ensure that Docker and Docker Compose are installed, then use the following command to build and start the application:

```bash
docker-compose up --build
```
This will set up the necessary services, including the PostgreSQL database and the FileStore application.

2. **Stop the Application**

To stop the application and remove containers, networks, and volumes created by Docker Compose, run:

```bash
docker-compose down
```

### API Endpoints

* Upload File

    *Endpoint* : POST /upload

    *Description*: Uploads a file and returns a unique file ID.

    *Request*:

    ```bash
    curl -F "file=@/path/to/your/file" http://localhost:8080/upload
    ```

    *Response*:

    ```json
    {
    "file_id": "unique-file-id"
    }
    ```

* Get Files Metadata

    *Endpoint*: GET /files

    *Description*: Retrieves metadata for all uploaded files.

    *Response*:

    ```json
    [
    {
        "id": "unique-file-id",
        "file_name": "file-name",
        "file_size": 12345
    }
    ]
    ```

* Download File

    *Endpoint*: GET /download

    *Query Parameters*:

        id: The unique file ID.
        Description: Downloads a file by its ID.

    *Request*:

    ```bash
    curl http://localhost:8080/download?id=unique-file-id --output file
    ```

### Testing
To run tests, use the following command:

```bash
go test ./...
```

This will run all unit, integration, and end-to-end tests in the tests directory.

### Configuration
Database connection settings and other configurations are managed through environment variables and the config package. Make sure to configure these settings before running the application.

### Contributing
Feel free to submit issues, suggest improvements, or contribute code via pull requests. Please ensure that your contributions adhere to the project's coding standards and include appropriate tests.

### License
This project is licensed under the MIT License. See the LICENSE file for details.