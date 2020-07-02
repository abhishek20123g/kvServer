package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/abhishek20123g/kvServer"
	"github.com/abhishek20123g/kvServer/grpcImp"
	"github.com/abhishek20123g/kvServer/rpcImp"
)

func main() {
	var client kvServer.Client
	methodName := "rpc"
	// In case method used is specified using Arguments
	if len(os.Args) > 1 {
		methodName = os.Args[1]
	}

	// Type casting Server variable according to method specified.
	switch methodName {
	case "rpc":
		client = new(rpcImp.Client)
	case "grpc":
		client = new(grpcImp.Client)
	default:
		fmt.Fprintf(
			os.Stderr, "unknown method name is specified %s. Method name should be (rpc, grpc) only", methodName,
		)
		os.Exit(0)
	}

	fmt.Printf("Client is connected to %s \n", kvServer.ServerAddr)
	fmt.Println("We can talk to server with the help of input buffers")
	var read = bufio.NewReader(os.Stdin)
	// Reading continuously, until the quit message is being has been specified.
	readStdin: for {
		line, _ := read.ReadString('\n')
		line = line[:len(line) - 1]
		args := strings.Split(line, " ")

		err := client.Connect("tcp", kvServer.ServerAddr)
		if err != nil {
			err = errors.New(fmt.Sprintf("Connect: %s", err.Error()))
			break readStdin
		}

		// Execution on the basis of command.
		switch {
		case strings.ToLower(args[0]) == "quit":
			break readStdin
		case strings.ToLower(args[0]) == "put" && len(args) == 3:
			err = client.Put(args[1], args[2])
			if err != nil {
				err = errors.New(fmt.Sprintf("Put: %s", err.Error()))
			} else {
				fmt.Println("Updated KV Storage")
			}
		case strings.ToLower(args[0]) == "get" && len(args) == 2:
			var value string
			value, err = client.Get(args[1])
			if err != nil {
				err = errors.New(fmt.Sprintf("Get: %s", err.Error()))
			} else {
				fmt.Printf("(key, value): (%s, %s) \n", args[1], value)
			}
		default:
			fmt.Println("unknown message command")
		}
		_ = client.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s \n", err.Error())
		}
	}
}
