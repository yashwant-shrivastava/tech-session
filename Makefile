PROTO_DIR := .
PROTO_GEN_DIR := .

install_protoc:
	brew install protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

compile_protos:
	find $(PROTO_DIR) -name "*.proto" \
	-exec protoc --go_out=$(PROTO_GEN_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(PROTO_GEN_DIR) \
	--go-grpc_opt=paths=source_relative {} \;

