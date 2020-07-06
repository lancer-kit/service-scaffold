#!/bin/bash

VERSION=1.1.0
COMMIT=$(git rev-parse --short HEAD)
TAG=$(git rev-parse --abbrev-ref HEAD)
SERVICE_PATH=lancer-kit/service-scaffold

if ! git diff-index --quiet --cached HEAD; then
  COMMIT=$COMMIT-dirty.
fi

go build -ldflags "-X $SERVICE_PATH/config.version=$VERSION -X $SERVICE_PATH/config.build=$COMMIT -X $SERVICE_PATH/config.tag=$TAG" .
