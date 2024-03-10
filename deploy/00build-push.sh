#!/bin/bash

docker build --target production -t harbor.deagwon.com/rbac-go/rbac-go:latest ../

docker push harbor.deagwon.com/rbac-go/rbac-go:latest

docker rmi harbor.deagwon.com/rbac-go/rbac-go:latest