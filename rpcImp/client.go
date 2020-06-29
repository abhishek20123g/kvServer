package rpcImp

import (
	"errors"
	"github.com/abhishek20123g/kvServer"
	"net/rpc"
)

// Client is the implementation of the client's data structure
// in case of rpc.
type Client struct {
	conn *rpc.Client
}

// Connect connects the client to the given network and address.
func (c *Client) Connect(network, address string) error {
	var err error
	c.conn, err = rpc.Dial(network, address)
	return err
}

// Put Update the (key, value) pair in KV Storage.
// Note: Before using Put make sure client is connected to
// the server.
func (c *Client) Put(key, value string) error {
	if c.conn == nil {
		return kvServer.ClientNotConnected
	}
	putArgs := &PutArgs{Key: key, Message: "Update KV Storage", Value: value}
	putReplyArgs := &GetReply{}
	return c.conn.Call("RpcServices.Put", putArgs, putReplyArgs)
}

// Get extract the value corresponds to given key from KV
// Storage.
// Note: Before using Get make sure client is connected to
// the server.
func (c *Client) Get(key string) (string, error) {
	if c.conn == nil {
		return "", kvServer.ClientNotConnected
	}
	getArgs := &GetArgs{Key: key, Message: "Requests data"}
	getReplyArgs := &GetReply{}
	err := c.conn.Call("RpcServices.Get", getArgs, getReplyArgs)
	if err != nil {
		return "", err
	}
	if getReplyArgs.Err != "" {
		err = errors.New(getReplyArgs.Err)
	}
	return 	getReplyArgs.Value, err
}

// Close closes the client connection to the server.
func (c *Client)Close() error {
	err := c.conn.Close()
	c.conn = nil
	return err
}
