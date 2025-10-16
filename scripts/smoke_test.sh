#!/usr/bin/env bash

kubectl apply -f k8s/install.yaml

# Wait for deployment to be ready
kubectl wait --for=condition=available --timeout=120s --namespace kosli deployment/hello-kosli

kubectl port-forward --namespace kosli svc/hello-kosli 8080:8080 &
sleep 5

ACTUAL_SHA=$(curl --retry 5 -s -X POST 'http://localhost:8080/version' | jq .git_sha)
CURRENT_SHA=\"$(git rev-parse HEAD)\"

if [ "$ACTUAL_SHA" != "$CURRENT_SHA" ]; then
  echo "Expected response to contain $CURRENT_SHA but got $ACTUAL_SHA"
  exit 1
fi

echo "Smoke test passed"