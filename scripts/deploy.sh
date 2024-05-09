#!/bin/bash

TIMER_IMAGE_NAME=timer-app-image
SCRAPER_IMAGE_NAME=scraper-app-image

docker build -t $TIMER_IMAGE_NAME -f ./docker/DockerfileTimer .
docker build -t $SCRAPER_IMAGE_NAME -f ./docker/DockerfileScraper .

kubectl apply -f ./k8s/timer-app-deployment.yaml --validate=true
kubectl apply -f ./k8s/timer-app-service.yaml --validate=true
kubectl apply -f ./k8s/scraper-app-deployment.yaml --validate=true

echo "Deployments and service applied"