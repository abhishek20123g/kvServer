## About
This project is for learning purposes. </br>
This project is simple implementation of `RPC API's` using `rpc` and `grpc` library separately.
A simple KV server is a simple map which allow us to update and read various value 
for the storage, stored in the server. 

#### Later
Current API's are designed in such a manner one can update KV Storage in the server only.

Update API's is such that each client can have there local Storage and one client
request data and update from other client's Storage.   

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

### Usage for protoc
```bash
# _GRPC_DST_DIR_ destination path for the service file
# _BUFFER_DST_DIR_ destination path for the proto file
# _IMPORT_PATH_ where we look for other import proto files.
protoc --proto_path=_IMPORT_PATH_ --go-grpc_out=_GRPC_DST_DIR_ --go_out=_BUFFER_DST_DIR_ _PROTO_SRC_PATH_/*.proto 
```

### Project Installation
```
# Getting project locally
go get 

cd ~/go/src/abhishek20123g/kvServer
make build
```

To start server and clients executable with rpc
```
./bin/server.o rpc
# OR
./bin/server.o
```
separate command
```
./bin/client.o rpc
# OR
./bin/client.o
```

To start server executable with grpc
```
./bin/server.o grpc
```
separate command
```
./bin/client.o grpc
```

To start node executable with grpc
```
./bin/node.o grpc
# OR
./bin/node.o rpc
```
