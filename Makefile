protoc -I api/proto/v1 \
-I${GOPATH}/src \
--go_out=plugins=grpc:pkg/api/v1 \
api/proto/v1/todo-service.proto