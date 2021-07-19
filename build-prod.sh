#!/bin/bash

docker build -t microservices ../
docker inspect microservices

# push
docker tag microservices gcr.io/sublime-etching-319904/microservices
docker push gcr.io/sublime-etching-319904/microservices