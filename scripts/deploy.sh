#!/bin/bash

GOLANG_IMAGE_NAME=server-app
PYTHON_IMAGE_NAME=scraper-app

docker build -t $GOLANG_IMAGE_NAME -f ./docker/DockerfileServer .
docker build -t $PYTHON_IMAGE_NAME -f ./docker/DockerfileScraper .

kubectl apply -f ./k8s/server-app-deployment.yaml --validate=true
kubectl apply -f ./k8s/server-app-service.yaml --validate=true
kubectl apply -f ./k8s/scraper-app-deployment.yaml --validate=true

echo "Deployments and service applied"