generate_grpc:
	protoc --go_out=. --go-grpc_out=. proto/energy.proto
	PATH=$PATH:$(pwd)/node_modules/.bin \
  protoc -I . \
  --es_out src/gen \
  --es_opt target=ts \
  a.proto b.proto c.proto

  protoc -I=. energy.proto \
    --js_out=import_style=commonjs:. \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:.

  protoc -I=. energy.proto \
  --js_out=import_style=commonjs,binary:$OUT_DIR \
  --grpc-web_out=import_style=typescript,mode=grpcweb:$OUT_DIR