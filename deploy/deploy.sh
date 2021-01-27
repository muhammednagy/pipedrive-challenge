#!/bin/bash

if [ $# -lt 3 ]; then
         echo -e "\nUsage:\n bash deploy.sh [project ID] [GKE cluster name] [GKE zone] [serviceaccount.json file path] \n"
         exit 1
fi

PROJECT_NAME=$1
GKE_CLUSTER=$2
GKE_ZONE=$3
SERVICE_ACCOUNT=$4
gcloud container clusters get-credentials "$GKE_CLUSTER" --zone "$GKE_ZONE"
if [ "$SERVICE_ACCOUNT" != ''  ]; then
  gcloud  auth activate-service-account --key-file="$SERVICE_ACCOUNT"
fi
gcloud --quiet config set project "$PROJECT_NAME"

sed -i "s/COMMIT_SHA/$GITHUB_SHA/g" ./app-deployment.yaml
kubectl apply -f dbdata-pvc.yaml || true # it's okay if the pvc is already there in case it was previously deployed
kubectl apply -f db-service.yaml
kubectl apply -f db-deployment.yaml
kubectl apply -f app-service.yaml
kubectl apply -f app-deployment.yaml

echo "Your app is now running at $(kubectl get svc app -o json | jq -r .status.loadBalancer.ingress[].ip)"