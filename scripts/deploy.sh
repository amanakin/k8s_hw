#!/bin/bash

TIMER_IMAGE_NAME=timer-app-image:v2.1
SCRAPER_IMAGE_NAME=scraper-app-image:v2.1

eval "$(minikube docker-env)"

docker build -t $TIMER_IMAGE_NAME -f ./docker/DockerfileTimer .
docker build -t $SCRAPER_IMAGE_NAME -f ./docker/DockerfileScraper .

istioctl install -y --set profile=demo --set meshConfig.outboundTrafficPolicy.mode=REGISTRY_ONLY
kubectl label namespace default istio-injection=enabled

kubectl apply -f ./k8s/timer-app-deployment.yaml --validate=true
kubectl apply -f ./k8s/timer-app-service.yaml --validate=true
kubectl apply -f ./k8s/scraper-app-deployment.yaml --validate=true

kubectl apply -f ./k8s/virtual-service.yaml --validate=true
kubectl apply -f ./k8s/ingress-gateway.yaml --validate=true
kubectl apply -f ./k8s/service-entry.yaml --validate=true

echo "Deployments and service applied"

minikube tunnel
