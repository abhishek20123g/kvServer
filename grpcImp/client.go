package grpcImp

import (
	"context"
	"errors"
	"time"

	"github.com/abhishek20123g/kvServer"
	"google.golang.org/grpc"
)

// Client is the implementation of the client's data structure
// in case of rpc.
type Client struct {
	conn *grpc.ClientConn
	client KVClient
}

// Connect connects the client to the given network and address.
func (c *Client) Connect(network, address string) error {
	var err error
	c.conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	c.client = NewKVClient(c.conn)
	return nil
}

// Put Update the (key, value) pair in KV Storage.
// Note: Before using Put make sure client is connected to
// the server.
func (c *Client) Put(key, value string) error {
	if c.conn == nil {
		return kvServer.ClientNotConnected
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	putArgs := &PutArgs{Key: key, Message: "Update KV Storage", Value: value}
	putReplyArgs, err := c.client.Put(ctx, putArgs)
	if err != nil {
		return err
	}
	if putReplyArgs.Err != "" {
		return errors.New(putReplyArgs.Err)
	}
	return nil
}

// Get extract the value corresponds to given key from KV Storage.
// Note: Before using Get make sure client is connected to
// the server.
func (c *Client) Get(key string) (string, error) {
	if c.conn == nil {
		return "", kvServer.ClientNotConnected
	}
	getArgs := &GetArgs{Key: key, Message: "Requests data"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getReply, err := c.client.Get(ctx, getArgs)
	if err != nil {
		return "", err
	}
	if getReply.Err != "" {
		return "", errors.New(getReply.Err)
	}
	return getReply.Value, nil
}

// Close closes the client connection to the server.
func (c *Client)Close() error {
	err := c.conn.Close()
	c.conn = nil
	return err
}
