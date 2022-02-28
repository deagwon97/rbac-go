#!/bin/bash

cd ~/src 

swag init

# convert Swagger docs to Open API docs
curl -X POST "https://converter.swagger.io/api/convert"\
    -H "accept: application/json" \
    -H "Content-Type: application/json" \
    -d "`cat docs/swagger.json`" > openapi.json