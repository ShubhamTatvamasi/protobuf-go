#!/bin/bash

# stop script if there are some errors
set -e

echo "Generating Proto Buffer code for golang..."

protoc -I pkg/ --go_out=pkg/ pkg/simple/simple.proto
protoc -I pkg/ --go_out=pkg/ pkg/enum/enum.proto
protoc -I pkg/ --go_out=pkg/ pkg/complex/complex.proto

echo "Go packages has been generated."
