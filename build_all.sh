#!/bin/bash

# Build each service
docker build -t api-gateway ./api-gateway
docker build -t users ./users
docker build -t products ./products
docker build -t orders ./orders
docker build -t payments ./payments