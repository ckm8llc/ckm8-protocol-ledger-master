#!/bin/bash

# Usage: 
#    integration/build/build.sh
#    integration/build/build.sh force # Always recreate docker image and container.
set -e

SCRIPTPATH=$(dirname "$0")

echo $SCRIPTPATH

if [ "$1" =  "force" ] || [[ "$(docker images -q ckm8_builder_image 2> /dev/null)" == "" ]]; then
    docker build -t ckm8_builder_image $SCRIPTPATH
fi

set +e
docker stop ckm8_builder
docker rm ckm8_builder
set -e

docker run --name ckm8_builder -it -v "$GOPATH:/go" ckm8_builder_image
