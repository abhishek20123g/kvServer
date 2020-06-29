// This file consist of all the datastructures common to
// the server and the client.
package rpcImp

//
// Put
//

// Put call arguments
type PutArgs struct {
	Key     string
	Value   string
	Message string
}

// Put reply arguments
type PutReply struct {
	Err string
}

//
// Get
//

// Get call arguments
type GetArgs struct {
	Key     string
	Message string
}

// Get reply arguments
type GetReply struct {
	Value string
	Err   string
}