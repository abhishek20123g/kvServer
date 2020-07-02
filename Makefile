SHELL=/bin/bash

# to see all colors, run
# bash -c 'for c in {0..255}; do tput setaf $c; tput setaf $c | cat -v; echo =$c; done'
# the first 15 entries are the 8-bit colors

# define standard colors
BLACK        := $(shell tput -Txterm setaf 0)
RED          := $(shell tput -Txterm setaf 1)
GREEN        := $(shell tput -Txterm setaf 2)
YELLOW       := $(shell tput -Txterm setaf 3)
LIGHTPURPLE  := $(shell tput -Txterm setaf 4)
PURPLE       := $(shell tput -Txterm setaf 5)
BLUE         := $(shell tput -Txterm setaf 6)
WHITE        := $(shell tput -Txterm setaf 7)

RESET := $(shell tput -Txterm sgr0)

# set target color
TARGET_COLOR := $(BLUE)

# Generate all the *.proto files in kvServer/grpcImp
generateproto:
	@echo "$(GREEN)[PROTO 1/1] Generating all proto and grpc_proto files in grpcImp $(RESET)"
	@protoc --go-grpc_out=./grpcImp --go_out=./grpcImp ./grpcImp/*.proto

# Build the client and server file.
build: clean generateproto
	@echo "$(GREEN)[BUILD 1/3] Building server executable $(type) $(RESET)"
	@go build -race -o ./bin/server.o ./main/server.go
	@echo "$(GREEN)[BUILD 2/3] Building client executable $(type) $(RESET)"
	@go build -race -o ./bin/client.o ./main/client.go
	@echo "$(GREEN)[BUILD 3/3] Building node executable $(type) $(RESET)"
	@go build -race -o ./bin/node.o ./main/node.go


# Clean all binary files related to the server and client.
# Clean all the proto generated files
clean:
	@echo "$(RED)[CLEAN 1/2] Cleaning all the binary files $(RESET)"
	@rm -rf bin/*.o
	@echo "$(RED)[CLEAN 2/2] Cleaning all the proto files $(RESET)"
	@rm -f ./grpcImp/*.pb.go