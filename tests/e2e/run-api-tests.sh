#!/usr/bin/env bash
set -x

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

APIURL=${APIURL:-http://localhost:8080}

npx newman run $SCRIPTDIR/AnimalRescue.postman_collection.json \
  --delay-request 500 \
  --global-var "APIURL=$APIURL"
