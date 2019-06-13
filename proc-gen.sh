#!/bin/bash

protoc -I api/proto/v1 \
-I third_party \
--go_out=plugins=grpc:pkg/api/v1 \
api/proto/v1/todo-service.proto


protoc -I api/proto/v1 \
       -I third_party \
       --grpc-gateway_out=logtostderr=true:pkg/api/v1 \
       api/proto/v1/todo-service.proto

protoc -I api/proto/v1 \
       -I third_party \
       --swagger_out=logtostderr=true:api/swagger/v1 \
       api/proto/v1/todo-service.proto
