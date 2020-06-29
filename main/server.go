package main

import (
	"fmt"
	"os"

	"github.com/abhishek20123g/kvServer"
	"github.com/abhishek20123g/kvServer/grpcImp"
	"github.com/abhishek20123g/kvServer/rpcImp"
)

func main() {
	var server kvServer.Server
	methodName := "rpc"
	// In case method used is specified using Arguments
	if len(os.Args) > 1 {
		methodName = os.Args[1]
	}

	// Type casting Server variable according to method specified.
	switch methodName {
	case "rpc":
		server = new(rpcImp.Server)
	case "grpc":
		server = new(grpcImp.Server)
	default:
		fmt.Fprintf(
			os.Stderr, "unknown method name is specified %s. Method name should be (rpc, grpc) only", methodName,
			)
		os.Exit(0)
	}

	var storage = kvServer.CreateNewKVStorage()
	err := server.InitServer("tcp", kvServer.ServerAddr, storage)
	defer func() {
		err = server.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error in closing the Server: %s", err.Error())
			os.Exit(0)
		}
	}()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error while strating the Server: %s", err.Error())
		os.Exit(0)
	}
}
