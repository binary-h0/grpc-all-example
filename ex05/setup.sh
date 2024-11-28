go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
cd proto
protoc -I. -I../../googleapis \
       --go_out=. --go-grpc_out=. \
       --grpc-gateway_out=. \
       user.proto