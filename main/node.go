// This example implementation focus to show the simultaneous behaviour
// as a Server and a Client. Each node have a specific port number
// to it. A node can request data from other node as well as fulfill
// the request of data from it self to other nodes.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/abhishek20123g/kvServer"
	"github.com/abhishek20123g/kvServer/grpcImp"
	"github.com/abhishek20123g/kvServer/rpcImp"
)

// This is just the reimplementation of the `server.go`
// that is Server handles the server side for the node.
// That is this function handles various request. from
// the other nodes.
func Server(server *kvServer.Server, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	var storage = kvServer.CreateNewKVStorage()
	err := (*server).InitServer("tcp", "127.0.0.1:" + port, storage)
	defer func() {
		err = (*server).Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error in closing the Server: %s \n", err.Error())
			return
		}
	}()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error while strating the Server: %s \n", err.Error())
		return
	}
}
// This is just the reimplementation of the `client.go`
// that is Client handles the server side for the node.
// That is this function handles various request. from
// the other nodes.
func Client(client *kvServer.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("We can talk to server with the help of input buffers")
	var read = bufio.NewReader(os.Stdin)
	// Reading continuously, until the quit message is being has been specified.
	readStdin: for {
		line, _ := read.ReadString('\n')
		line = line[:len(line) - 1]
		args := strings.Split(line, " ")
		port := args[0]
		err := (*client).Connect("tcp", "127.0.0.1:" + port)
		if err != nil {
			err = errors.New(fmt.Sprintf("Connect: %s", err.Error()))
			break readStdin
		}
		fmt.Printf("Client is connected to 127.0.0.1:%s \n", port)

		// Execution on the basis of command.
		switch {
		case strings.ToLower(args[1]) == "quit":
			break readStdin
		case strings.ToLower(args[1]) == "put" && len(args) == 4:
			err = (*client).Put(args[2], args[3])
			if err != nil {
				err = errors.New(fmt.Sprintf("Put: %s", err.Error()))
			} else {
				fmt.Printf("Updated KV Storage of port %s \n", port)
			}
		case strings.ToLower(args[1]) == "get" && len(args) == 3:
			var value string
			value, err = (*client).Get(args[2])
			if err != nil {
				err = errors.New(fmt.Sprintf("Get: %s", err.Error()))
			} else {
				fmt.Printf("(key, value): (%s, %s) for port %s \n", args[2], value, port)
			}
		default:
			fmt.Println("unknown message command")
		}
		_ = (*client).Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s \n", err.Error())
		}
	}
}

func main() {
	port := "8000"
	methodName := "rpc"
	var wg sync.WaitGroup
	var server kvServer.Server
	var client kvServer.Client

	// Read various os arguments.
	switch len(os.Args) {
	// When only method name is specified.
	case 2:
		methodName = os.Args[1]
	// When both method name and port is specified.
	case 3:
		methodName = os.Args[1]
		port = os.Args[2]
	default:
		fmt.Fprintf(os.Stderr, "command line arguments has not been specified properly")
	}

	// Type casting Server variable according to method specified.
	switch methodName {
	case "rpc":
		server = new(rpcImp.Server)
		client = new(rpcImp.Client)
	case "grpc":
		server = new(grpcImp.Server)
		client = new(grpcImp.Client)
	default:
		fmt.Fprintf(
			os.Stderr, "unknown method name is specified %s. Method name should be (rpc, grpc) only \n", methodName,
		)
		os.Exit(0)
	}

	// Start various subroutine.
	wg.Add(2)
	go Server(&server, port, &wg)
	go Client(&client, &wg)

	// Wait till all the subroutine hasn't been finished.
	wg.Wait()
}