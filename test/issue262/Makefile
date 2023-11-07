regenerate:
	go install github.com/coderyw/protobuf/protoc-min-version
	go install github.com/coderyw/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/coderyw/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/coderyw/protobuf/types:. timefail.proto
