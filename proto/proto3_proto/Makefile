regenerate:
	go install github.com/coderyw/protobuf/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mtest_proto/test.proto=github.com/coderyw/protobuf/proto/test_proto,\
	Mgoogle/protobuf/any.proto=github.com/coderyw/protobuf/types:. \
	--proto_path=../../protobuf:../:. proto3.proto
	
