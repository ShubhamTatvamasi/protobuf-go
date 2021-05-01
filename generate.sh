#!/bin/bash

echo "Hello"
protoc -I pkg/ --go_out=pkg/ pkg/simple/simple.proto 



