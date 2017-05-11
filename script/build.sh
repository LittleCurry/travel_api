#!/bin/bash

function build_docker {
    TAG=$1

    IMAGE_NAME="registry.iguiyu.com/park/api"

    IMAGE_FULL_NAME="$IMAGE_NAME:$TAG"
    HAS_OLD_IMAGES=$(docker images|grep $IMAGE_NAME|grep $TAG|wc -l)
    echo $HAS_OLD_IMAGES
    if [ $HAS_OLD_IMAGES -ne "0" ]; then
        echo "Remove docker image..."
        docker rmi $IMAGE_FULL_NAME
    fi

    echo "Building docker image..."
    docker build -t $IMAGE_FULL_NAME .

    echo "Push image to reigstry"
    docker push $IMAGE_FULL_NAME
}

set -e

DATETAG=$(date +"%y%m%d%H%M%S")

cd $GOPATH/src/git.iguiyu.com/park/api

echo "Building application..."
CGO_ENABLED=0 GOOS=linux go build -o resource/main .

build_docker $DATETAG

build_docker "latest"

echo "Cleanup resources..."
rm resource/main

echo "Done"
