#!/bin/bash

# Generate Go code
protoc --go_out=../api_service --go-grpc_out=../api_service ml.proto

# Generate Python code
python -m grpc_tools.protoc -I. --python_out=../ml_service/services  --grpc_python_out=../ml_service/services ml.proto

