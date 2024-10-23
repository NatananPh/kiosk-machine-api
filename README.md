# kiosk-machine-api

## Introduction
The project implements a REST API for a kiosk machine using the Echo framework in Go, connected to a PostgreSQL database for reliable data management. PostgreSQL was chosen for its robustness and ACID compliance, ensuring data consistency. The project follows Go Clean Architecture for maintainability and scalability. It includes an authentication system with role-based access control (admin, user) through middleware to restrict access to admin APIs. Secure authentication is achieved using JWT, and passwords are hashed and salted to protect sensitive information.

## Setup steps
### Start PostgreSQL on Docker
1. Run ```go mod tidy``` for Installing dependencies
   
   ```
   go mod tidy
   ```
   
3. Run docker compose-command below
   
   ```
   docker-compose -f .\docker-compose.yml up -d
   ```
4. Run migration script and mocking data for testing
   
   ```
   go run .\database\migrationV2\migrationV2.go
   ```
5. Start server using ```air```
   ```
   air
   ```
## Examples of API Usage
### 1. Login API
- **Endpoint:** `localhost:8080/v1/auth/login
- **Method:** `POST`
- **Request Body for admin (success):**
  ```json
  {
    "username": "admin",
    "password": "admin"
  }
  ```
- **Request Body for user (success):**
  ```json
  {
    "username": "user",
    "password": "user"
  }
  ```
- **Response (success): we will use token for ensuring that users cannot access admin API**
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5...."
  }
  ```
- **Response (failed)**
  ```json
  {
     "error": "invalid username or password"
  }
  ```
### 2. Get Products API
- **Endpoint:** `localhost:8080/v1/products
- **Method:** `GET`
- We can add query parameters (category, page, linit) for filter and pagination
- **Example Endpoint:** ```localhost:8080/v1/products?category=Snack&page=1&limit=1```
  - **Response (success)**
  ```json
  [
    {
        "id": 5,
        "name": "Lay",
        "price": 20,
        "amount": 10,
        "category": "Snack"
    }
   ]
  ```
### 3. Get Product by id API
### 4. Create product API
### 5. Update product by id API
### 6. Delete product by id API
### 7. Purchase product API

