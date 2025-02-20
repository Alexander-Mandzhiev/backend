package grpc_client

import (
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClientManager struct {
	connections map[string]*grpc.ClientConn
	mu          sync.Mutex
}

func NewGRPCClientManager() *GRPCClientManager {
	return &GRPCClientManager{
		connections: make(map[string]*grpc.ClientConn),
	}
}

func (m *GRPCClientManager) GetClientConn(addr string) (*grpc.ClientConn, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if conn, exists := m.connections[addr]; exists {
		return conn, nil
	}

	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(16<<20)))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to service at %s: %w", addr, err)
	}

	m.connections[addr] = conn
	return conn, nil
}

func (m *GRPCClientManager) CloseAll() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for addr, conn := range m.connections {
		if err := conn.Close(); err != nil {
			fmt.Printf("Failed to close connection to %s: %v\n", addr, err)
		}
	}
	m.connections = make(map[string]*grpc.ClientConn)
}
