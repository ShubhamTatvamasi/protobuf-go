#!/bin/bash

# stop script if there are some errors
set -e

echo "Generating Proto Buffer code for golang..."
protoc -I pkg/ --go_out=pkg/ pkg/simple/simple.proto 

echo "Done"
