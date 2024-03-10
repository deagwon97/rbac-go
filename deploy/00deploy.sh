#!/bin/bash

./00build-push.sh

kubectl rollout restart deployment/rbac-go -n rbac-go