generate_grpc:
	protoc --go_out=. --go-grpc_out=. proto/energy.proto