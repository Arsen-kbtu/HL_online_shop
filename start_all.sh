#!/bin/bash

# Run each service
docker run -d -p 8080:8080 --name api-gateway api-gateway
docker run -d -p 8081:8081 --name users users
docker run -d -p 8082:8082 --name products products
docker run -d -p 8083:8083 --name orders orders
docker run -d -p 8084:8084 --name payments payments
