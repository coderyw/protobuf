regenerate:
	go install github.com/coderyw/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/coderyw/protobuf/protobuf/ \
	--gogo_out=. events.proto 
