#!/bin/bash

# Install microk8s
sudo snap install microk8s --classic

# Wait for microk8s to become ready
microk8s status --wait-ready

# create alias

alias kubectl='microk8s kubectl'

# Enable the microk8s registry
microk8s enable registry

# Build the Docker image
docker build . -t localhost:32000/server:latest

# Show the Docker images
docker images

# Get the image ID and tag it
IMAGE_ID=$(docker images -q localhost:32000/server:latest)
docker tag $IMAGE_ID localhost:32000/server:latest

# Push the Docker image to the registry
docker push localhost:32000/server

#kubernetes command to run everything

kubectl apply -f kubernetes-replicaset.yaml
