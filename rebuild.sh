#!/bin/bash

docker container stop pdf-server && docker container rm pdf-server
docker build -t pdf-server . --no-cache --progress=plain --debug
docker run --name pdf-server -d -p 8080:80 pdf-server
