# kiosk-machine-api

## Introduction

The project implements a REST API for a kiosk machine using the Echo framework in Go, connected to a PostgreSQL database for reliable data management. PostgreSQL was chosen for its robustness and ACID compliance, ensuring data consistency. The project follows Go Clean Architecture for maintainability and scalability. It includes an authentication system with role-based access control (admin, user) through middleware to restrict access to admin APIs. Secure authentication is achieved using JWT, and passwords are hashed and salted to protect sensitive information.

## Setup steps

### Start PostgreSQL on Docker

1. Run `go mod tidy` for Installing dependencies

   ```
   go mod tidy
   ```

2. Run docker compose-command below

   ```
   docker-compose -f .\docker-compose.yml up -d
   ```

3. Run migration script and mocking data for testing

   ```
   go run .\database\migrationV2\migrationV2.go
   ```

4. Start server using `air`
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
- **Response (success) 200: we will use token for ensuring that users cannot access admin API**
  ![image](https://github.com/user-attachments/assets/a58313cb-5ecb-4143-a4d7-33624cfbaeb2)

  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5...."
  }
  ```

- **Response (failed) 400**
  ```json
  {
    "error": "invalid username or password"
  }
  ```

### 2. Get Products API

- **Endpoint:** `localhost:8080/v1/products
- **Method:** `GET`
- We can add query parameters (category, page, linit) for filter and pagination
- **Example Endpoint:** `localhost:8080/v1/products?category=Snack&page=1&limit=1`
  - **Response (success) 200**
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

- **Endpoint:** `localhost:8080/v1/products/{id}
- **Method:** `GET`
- **Response (success) 200**
  ```json
  {
    "id": 3,
    "name": "Fanta",
    "price": 20,
    "amount": 10,
    "category": "Drink"
  }
  ```
- **Response (failed) 404: product not found**
  ```json
  {
    "error": "Product not found"
  }
  ```

### 4. Create product API

- **Endpoint:** `localhost:8080/v1/products
- **Method:** `POST`
- **Request body**
  ```json
  {
    "name": "Testo",
    "price": 25,
    "amount": 100,
    "category": "Snack"
  }
  ```
- **Response (success) 201**
  ```json
  {
    "id": 8,
    "name": "Testo",
    "price": 25,
    "amount": 100,
    "category": "Snack"
  }
  ```
- **Response (failed) 401: Invalid token**
  ```json
  {
    "message": "Unauthorized: Invalid token"
  }
  ```
- **Response (failed) 400: Invalid request e.g. when price or amount is negative number**
  ```json
  {
    "error": "Invalid request"
  }
  ```
- **Response (failed) 403: admin access required**
  ```json
  {
    "message": "Forbidden: Admin access required"
  }
  ```

### 5. Update product by id API

- **Endpoint:** `localhost:8080/v1/products/{id}
- **Method:** `PUT`
- **Request body**
  ```json
  {
    "name": "Spritess",
    "price": 1000,
    "amount": 10,
    "category": "Drinkkk"
  }
  ```
- **Response (success) 200**
  ```json
  {
    "message": "Product updated"
  }
  ```
- **Response (failed) 404: product not found**
  ```json
  {
    "error": "Product not found"
  }
  ```
  - **Response (failed) 400: Invalid request e.g. id is not number (`/v1/products/abc`)**
  ```json
  {
    "error": "Invalid request"
  }
  ```
  - **Response (failed) 403: admin access required**
  ```json
  {
    "message": "Forbidden: Admin access required"
  }
  ```

### 6. Delete product by id API

- **Endpoint:** `localhost:8080/v1/products/{id}
- **Method:** `DELETE`
- **Response (success) 200**
  ```json
  {
    "message": "Product deleted"
  }
  ```
- **Response (failed) 404: product not found**
  ```json
  {
    "error": "Product not found"
  }
  ```
  - **Response (failed) 400: Invalid request e.g. id is not number (`/v1/products/abc`)**
  ```json
  {
    "error": "Invalid request"
  }
  ```
  - **Response (failed) 403: admin access required**
  ```json
  {
    "message": "Forbidden: Admin access required"
  }
  ```

### 7. Purchase product API

- **Endpoint:** `localhost:8080/v1/products/{id}/purchase
- **Method:** `POST
- **Request Body**
  ```json
  {
    "payment_amount": 100
  }
  ```
- **Response (success) 200**
  ```json
  {
    "product_id": 2,
    "change": {
      "10": 1,
      "20": 1,
      "50": 1
    }
  }
  ```
- **Response (failed) 404: product not found**
  ```json
  {
    "error": "Product not found"
  }
  ```
  - **Response (failed) 400: Invalid request**
  ```json
  {
    "error": "Invalid request"
  }
  ```
  - **Response (failed) 400: Invalid request (Insufficient money)**
  ```json
  {
    "error": "Insufficient money"
  }
  ```
  - **Response (failed) 400: Invalid request (product out of stock)**
  ```json
  {
    "error": "Product out of stock"
  }
  ```

## Postman Collection and ENV

You can download the Postman collection and environment file from the following Google Drive link:

[Postman Collection and ENV](https://drive.google.com/file/d/1tOsTMOVLNfz5dhFCdQWuB6ANDVLuCnvD/view?usp=sharing)
