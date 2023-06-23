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

   

