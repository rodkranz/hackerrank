#!/usr/bin/env bash

output=$(echo -e "cde\nabc" | go run main.go)
echo "Must be [4 = ${output}]"

output=$(echo -e "ilovea\nchallenge" | go run main.go)
echo "Must be [9 = ${output}]"
