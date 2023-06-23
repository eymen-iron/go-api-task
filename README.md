# go-api-task

# API Task

This project implements an API task with CRUD operations using the Go programming language. It was initially requested to be developed in PHP, but it has been converted to Go for implementation.

## Requirements

- Go 1.19 or higher

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/eymen-iron/go-api-task.git
    ```
   
2. Change directory to the project folder:

   ```bash
   cd go-api-task
   ```
3. Install the necessary dependencies:

   ```bash
   go mod download
   ```
   
4. Build the project:

   ```bash
    go build -o go-api-task .
    ```
   
## Usage

1. Run the project:

   ```bash
   ./go-api-task
   ```
   
2. The API will be accessible at `http://localhost:3005`.
3. You can use the following endpoints to interact with the API:
    - `GET /` - Get all structures
    - `GET /{id}` - Get a structure by ID
    - `CATCH /{id}` - Update a structure by ID
    - `DELETE /{id}` - Update a structure status by ID 

## Migration from PHP to Go

The original API task was developed in PHP, but it has been migrated to Go for implementation. The conversion process involved the following steps:

1. Rewriting the API endpoints in Go using the Go Fiber framework.
2. Updating the database access code to use the Go SQL package and SQLite.
3. Converting PHP-specific code to equivalent Go code.
4. Handling error cases and updating error handling mechanisms.
5. Rebuilding and testing the application in Go.

Please refer to the codebase for the complete Go implementation of the API task.



