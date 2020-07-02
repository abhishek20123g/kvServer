package grpcImp

import (
	"context"
	"fmt"
	"net"

	"github.com/abhishek20123g/kvServer"
	"google.golang.org/grpc"
)

// grpcServices Wrapper over kvServer.KVStorage which has updated method
// according to rpc service.
type GrpcServices struct {
	UnimplementedKVServer
	data *kvServer.KVStorage
}

// Implementation of Put service.
func (p *GrpcServices) Put(ctx context.Context, args *PutArgs) (*PutReply, error) {
	var reply = new(PutReply)
	err := p.data.Put(args.Key, args.Value)
	if err != nil {
		reply.Err = err.Error()
	}
	return reply, nil
}

// Implementation of the Get service.
func (p *GrpcServices) Get(ctx context.Context, args *GetArgs) (*GetReply, error) {
	fmt.Println(args.Message)
	var err error
	var reply = new(GetReply)
	reply.Value, err = p.data.Get(args.Key)
	if err != nil {
		reply.Err = err.Error()
	}
	return reply, nil
}

// Server is structure to handle the server in case of grpc.
type Server struct {
	serviceHandler *GrpcServices
	server         *grpc.Server
}

// InitServer initialises the Server connection.
// NOTE: Ideally it should only initialise the variable not
// run or close the server connection, But in this case it
// was easy to code this way.
func (s *Server) InitServer(network, address string, data *kvServer.KVStorage) error {
	// Created a service handler for the Server.
	s.serviceHandler = new(GrpcServices)
	s.serviceHandler.data = data
	s.server = grpc.NewServer()
	RegisterKVServer(s.server, s.serviceHandler)

	// listen for the new connections.
	listen, err := net.Listen(network, address)
	if err != nil {
		return err
	}

	fmt.Printf("Server has Initialised at %s. Started listen for the new connections. \n", address)
	if err := s.server.Serve(listen); err != nil {
		return err
	}
	return nil
}

// Close closed the server connection.
// NOTE: Unimplemented Close function as connection is closed
// by the InitServer().
func (s *Server) Close() error {
	return nil
}
