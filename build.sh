#!/bin/bash

docker build -t gofrit .
docker run -it --rm -p 8080:8080 --name gofritserv gofrit