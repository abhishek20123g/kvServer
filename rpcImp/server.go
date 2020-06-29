package rpcImp

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/abhishek20123g/kvServer"
)

// RpcServices Wrapper over kvServer.KVStorage which has updated method
// according to rpc service.
type RpcServices struct {
	data *kvServer.KVStorage
}

// Implementation of the Get service.
func (p *RpcServices) Get(args *GetArgs, reply *GetReply) error {
	fmt.Println(args.Message)
	var err error
	reply.Value, err = p.data.Get(args.Key)
	if err != nil {
		reply.Err = err.Error()
	}
	return nil
}

// Implementation of Put service.
func (p *RpcServices) Put(args *PutArgs, reply *PutReply) error {
	err := p.data.Put(args.Key, args.Value)
	if err != nil {
		reply.Err = err.Error()
	}
	return nil
}

// Server is structure to handle the server in case of rpc.
type Server struct {
	serviceHandler *RpcServices
	server         *rpc.Server
}

// InitServer initialises the Server connection.
// NOTE: Ideally it should only initialise the variable not
// run or close the server connection, But in this case it
// was easy to code this way.
func (s *Server) InitServer(network, address string, data *kvServer.KVStorage) error {
	// Created a service handler for the Server.
	s.serviceHandler = new(RpcServices)
	s.serviceHandler.data = data
	s.server = rpc.NewServer()
	err := s.server.Register(s.serviceHandler)
	if err != nil {
		return err
	}

	// listen for the new connections.
	listen, err := net.Listen(network, address)
	if err != nil {
		return err
	}

	fmt.Printf("Server has Initialised at %s. Started listen for the new connections. \n", address)
	// Initialise the never ending server.
	for {
		client, err := listen.Accept()

		if err != nil {
			return err
		}
		// Handle each client with separate thread.
		go s.server.ServeConn(client)
	}
}

// Close closed the server connection.
// NOTE: Unimplemented Close function as connection is closed
// by the InitServer().
func (s *Server) Close() error {
	return nil
}
