PROTO_DIR := protobufs/proto
PROTO_GEN_DIR := .

install_protoc:
	brew install protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

compile_protos:
	find $(PROTO_DIR) -name "*.proto" \
	-exec protoc --go_out=$(PROTO_GEN_DIR) \
	--go_opt=paths=source_relative {} \;














