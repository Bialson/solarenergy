generate_grpc:
	protoc --go_out=. --go-grpc_out=. energy.proto