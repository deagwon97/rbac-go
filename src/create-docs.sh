#!/bin/bash

cd ~/src 

swag init

curl -X POST "https://converter.swagger.io/api/convert"\
    -H "accept: application/json" \
    -H "Content-Type: application/json" \
    -d "`cat docs/swagger.json`" > openapi.json