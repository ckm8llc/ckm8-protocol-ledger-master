#!/bin/bash

# Build a docker image for a ckm8 node.
# Usage: 
#    integration/docker/node/build.sh
#
# After the image is built, you can create a container by:
#    docker stop ckm8_node
#    docker rm ckm8_node
#    docker run -e ckm8_CONFIG_PATH=/ckm8/integration/privatenet/node --name ckm8_node -it ckm8
set -e

SCRIPTPATH=$(dirname "$0")

echo $SCRIPTPATH

if [ "$1" =  "force" ] || [[ "$(docker images -q ckm8 2> /dev/null)" == "" ]]; then
    docker build -t ckm8 -f $SCRIPTPATH/Dockerfile .
fi


