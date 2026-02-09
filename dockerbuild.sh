#!/bin/bash
set -e

# -------- CONFIG --------
BIN_NAME="lmsapieng"
MAIN_PKG="./"
DOCKERFILE="Dockerfile"
# ------------------------

export DOCKER_BUILDKIT=1

docker build \
  -f ${DOCKERFILE} \
  --build-arg BIN_NAME=${BIN_NAME} \
  --build-arg MAIN_PKG=${MAIN_PKG} \
  --output type=local,dest=. \
  .

echo "Binary created: ./lmsapieng"
cp lmsapieng lmsapieng_4008