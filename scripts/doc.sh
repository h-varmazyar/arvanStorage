#!/bin/sh
echo "generating swagger API docs for ${1}"

swag fmt
swag init -g services/${1}/cmd/main.go -o ./services/${1}/docs