generate_grpc:
	protoc --go_out=. --go-grpc_out=. proto/energy.proto

  protoc -I=. energy.proto \
    --js_out=import_style=commonjs:. \
    --grpc-web_out=import_style=typescript,mode=grpcweb:.