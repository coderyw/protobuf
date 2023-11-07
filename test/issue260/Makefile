regenerate:
	go install github.com/coderyw/protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/coderyw/protobuf/types:. --proto_path=../../../../../:../../protobuf/:. issue260.proto
