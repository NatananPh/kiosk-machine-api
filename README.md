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
### 2. Get Products API
### 3. Get Product by id API
### 2. Create product API
### 2. Update product by id API
### 2. Delete product by id API
### 2. Purchase product API

