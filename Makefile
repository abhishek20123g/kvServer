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
	@echo "$(RED)Remove all proto and grpc files in grpcImp. $(RESET)"
	rm -rf ./grpcImp/*.pb.go
	@echo "$(BLUE)Generate updated proto and grpc file in grpcImp $(RESET)"
	protoc --go-grpc_out=./grpcImp --go_out=./grpcImp ./grpcImp/*.proto

build:
	@echo "$(RED)Clean all previous binaries. $(RESET)"
	rm -f *.o
	@echo "$(GREEN)build server file $(type) $(RESET)"
	go build -o server.o ./main/server.go $(type)
	@echo "$(GREEN)build the client file $(type) $(RESET)"
	go build -o client.o ./main/client.go $(type)

clean:
	@echo "$(RED)Clean all the binary files $(RESET)"
	rm -f *.o