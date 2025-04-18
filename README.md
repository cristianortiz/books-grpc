# Golang gRPC Backend

This project is a backend implementation in Go (Golang) that uses **gRPC** for communication. It provides a service for managing books, including operations such as adding, updating, retrieving, and deleting books. The project is designed with a modular structure, making it easy to extend and maintain.

---

## **Project Structure**

Below is the folder and file structure of the project:

books-grpc/
├── cmd/
│ └── rest-books-server/
│ └── main.go # Entry point for the REST API server
├── configs/
│ └── config.yaml # Configuration file for the application
├── internal/
│ ├── grpc-books-server/ # gRPC server implementation
│ │ └── grpc-books-server.go
│ ├── rest-books-server/ # REST API server implementation
│ │ ├── handlers.go # HTTP handlers for REST endpoints
│ │ ├── mapper.go # Utility for mapping between models
│ │ ├── rest-books-server.go # Main logic for the REST server
│ │ └── router.go # Router setup for REST endpoints
│ ├── pkg/
│ │ ├── configs/ # Configuration loading and management
│ │ │ └── configs.go
│ │ ├── db/ # Database connection and migrations
│ │ │ ├── db.go # Database connection setup
│ │ │ └── migrations/
│ │ │ └── migrator.go # Database migration logic
│ │ ├── model/ # Data models for business logic and database
│ │ │ ├── book.go # Business logic model for books
│ │ │ └── book-db.go # Database-specific model for books
│ │ ├── proto/ # gRPC protocol buffer definitions
│ │ │ ├── book.proto # Protocol buffer definition
│ │ │ ├── book.pb.go # Generated Go code from book.proto
│ │ │ └── book_grpc.pb.go # Generated gRPC server and client code
│ │ ├── repository/ # Data access layer
│ │ │ └── book-repo.go # Repository for book-related database operations
│ │ └── service/ # Business logic layer
│ │ └── book-service.go # Service for managing books
├── scripts/
│ ├── Dockerfile # Dockerfile for building the API
│ └── docker-compose.yml # Docker Compose file for API and database
├── go.mod # Go module dependencies
└── README.md # Project documentation

---

## **Key Components**

### **1. REST API Server**

- **Location**: `internal/rest-books-server/`
- **Description**: Implements a RESTful API for managing books. It uses the `gorilla/mux` router to define endpoints for CRUD operations.
- **Endpoints**:
  - `GET /books`: Retrieve a list of all books.
  - `POST /books`: Add a new book.
  - `PUT /books`: Update an existing book.
  - `GET /books/{isbn}`: Retrieve a book by its ISBN.
  - `DELETE /books/{isbn}`: Remove a book by its ISBN.

### **2. gRPC Server**

- **Location**: `internal/grpc-books-server/`
- **Description**: Implements a gRPC server for managing books. It uses protocol buffers (`proto`) to define the gRPC service and messages.
- **gRPC Methods**:
  - `AddBook`: Add a new book.
  - `ListBooks`: Retrieve all books.
  - `GetBook`: Retrieve a book by its ISBN.
  - `RemoveBook`: Remove a book by its ISBN.
  - `UpdateBook`: Update an existing book.

### **3. Configuration**

- **Location**: `configs/config.yaml` and `internal/pkg/configs/`
- **Description**: Handles application configuration, including database connection settings, server settings, and other parameters. The configuration is loaded using the `viper` library and supports environment variable overrides.

### **4. Database**

- **Location**: `internal/pkg/db/`
- **Description**: Manages the database connection and migrations.
  - **Connection**: Uses `gorm` to connect to a PostgreSQL database.
  - **Migrations**: Uses `golang-migrate` to handle database schema migrations.

### **5. Models**

- **Location**: `internal/pkg/model/`
- **Description**: Defines the data models used in the application.
  - `Book`: Represents the business logic model for books.
  - `DBBook`: Represents the database-specific model for books.

### **6. Repository**

- **Location**: `internal/pkg/repository/`
- **Description**: Implements the data access layer for interacting with the database. Provides CRUD operations for books.

### **7. Service**

- **Location**: `internal/pkg/service/`
- **Description**: Implements the business logic for managing books. It interacts with the repository layer to perform operations.

---

## **How to Run the Project**

### **1. Prerequisites**

- Install Docker and Docker Compose.
- Install Go (version 1.24 or higher).

### **2. Running with Docker Compose**

1. Build and start the services:
   ```bash
   docker-compose up --build
   ```

```

```
