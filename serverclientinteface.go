package kvServer

import (
	"errors"
	"sync"
)

// ServerAddr specify the address where the server will start.
const ServerAddr = "127.0.0.1:8000"
var KeyNotFound error = errors.New("key not found")
var ClientNotConnected = errors.New("client object is not connected to the Server")

// KVStorage manages the storage for the KV server.
// KVStorage abstract away the details.
type KVStorage struct {
	data  map[string]string
	mutex sync.Mutex
}

// CreateNewKVStorage create new object for the KVStorage.
func CreateNewKVStorage() *KVStorage {
	return &KVStorage{data: make(map[string]string)}
}

// Get provides the value for the given key from the storage.
func (kv *KVStorage) Get(key string) (string, error) {
	kv.mutex.Lock()
	defer kv.mutex.Unlock()

	if value, ok := kv.data[key];ok {
		return value, nil
	} else {
		return "", KeyNotFound
	}
}

// Put updates the (key, value) pair in the storage.
func (kv *KVStorage) Put(key, value string) error {
	//kv.mutex.Lock()
	//defer kv.mutex.Unlock()

	kv.data[key] = value
	return nil
}

// Common Server Interface to both rpcImp and grpcImp.
type Server interface {
	// Initialise the server variable at the given address
	// for the specific network.
	InitServer(network, address string, data *KVStorage) error
	// Close closes the Server.
	Close() error
}

// Common Client Interface to both rpcImp and grpcImp.
type Client interface {
	// Initialise the server variable at the given address
	// for the specific network.
	Connect(network, address string) error
	// Allow us to extract the data at the KVServer.
	Get(key string) (string, error)
	// Allow us to update the data at the KVServer.
	Put(key, value string) error
	// Close closes the client connection to the server.
	Close() error
}
