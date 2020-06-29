### Environment Setup
```bash
# Installation of protoc
PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'

# Download all the dependencies
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc

# Instal all the dependencies
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Add (go env GOPATH)/bin to PATH variable.
echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.bashrc
```

### Usage
```bash
# _GRPC_DST_DIR_ destination path for the service file
# _BUFFER_DST_DIR_ destination path for the proto file
# _IMPORT_PATH_ where we look for other import proto files.
protoc --proto_path=_IMPORT_PATH_ --go-grpc_out=_GRPC_DST_DIR_ --go_out=_BUFFER_DST_DIR_ _PROTO_SRC_PATH_/*.proto 
```