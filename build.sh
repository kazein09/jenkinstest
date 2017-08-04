#!/bin/sh

APP_NAME=main
EXTRACT_NAME=extract

# Build with go in Docker
docker build -t mygolang:build -f Dockerfile.build ./

# Extract the Go binary
docker create --name "${EXTRACT_NAME}" mygolang:build
docker cp "${EXTRACT_NAME}:/${APP_NAME}" "./${APP_NAME}"  
docker rm -f "${EXTRACT_NAME}"

# Package the application with Docker and the go binary
docker build --no-cache -t mygolang:latest ./
