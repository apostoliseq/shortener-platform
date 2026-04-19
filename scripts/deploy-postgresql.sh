#!/bin/bash
set -e

helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update

helm upgrade --install postgresql bitnami/postgresql \
  --namespace shortener \
  --create-namespace \
  --values helm/postgresql-values.yaml \
  --version 18.5.24