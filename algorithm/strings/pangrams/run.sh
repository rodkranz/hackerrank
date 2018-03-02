#!/usr/bin/env bash

output=$(echo -e "We promptly judged antique ivory buckles for the next prize" | go run main.go)
echo "Must be [pangram = ${output}]"
output=$(echo -e "We promptly judged antique ivory buckles for the prize" | go run main.go)
echo "Must be [no pangram = ${output}]"
