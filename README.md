# HL_online_shop

HL_online_shop is a microservices-based online store management system. The project includes microservices for Users, Products, Orders, Payments, and an API Gateway. The services are containerized using Docker and managed with Docker Compose. The deployment is handled on Render and DOES NOT WORK PROPERLY.
https://hl-online-shop.onrender.com/swagger/

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)


## Features

- User management
- Product management
- Order processing
- Payment processing
- API Gateway for routing and load balancing
- RESTful APIs with CRUD operations
- Swagger documentation for APIs
- Containerization with Docker
- Deployment with Render

## Technologies Used

- Go
- Gorilla Mux
- Gorm
- PostgreSQL
- Docker
- Docker Compose
- Render
- Swaggo (Swagger for Go)

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go 1.19 or later

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Arsen-kbtu/HL_online_shop.git
   cd HL_online_shop
   ```
2. Copy the .env.example to .env and update the environment variables as needed:
    ```bash
   cp .env.example .env
   ```
3. Build and start the services using Docker Compose:
    ```bash
    make build
    make up
    ```
## Running Services
Users service: http://localhost:8081  
Products service: http://localhost:8082  
Orders service: http://localhost:8083  
Payments service: http://localhost:8084  
API Gateway: http://localhost:8080  

## API Documentation
The project uses Swaggo to generate Swagger documentation. You can access the API documentation at:
http://localhost:8080/swagger/index.html
