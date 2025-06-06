package factory

import (
	"context"
	"fmt"
)

func (p *ClientProvider) getClient(ctx context.Context, serviceType ServiceType) (interface{}, error) {
	p.clientsMutex.RLock()
	client, exists := p.clients[serviceType]
	p.clientsMutex.RUnlock()

	if exists {
		return client, nil
	}

	addr, ok := p.serviceMap[serviceType]
	if !ok {
		return nil, fmt.Errorf("service %s not configured", serviceType)
	}

	conn, err := p.manager.GetClientConn(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", serviceType, err)
	}

	p.clientsMutex.Lock()
	defer p.clientsMutex.Unlock()

	if client, exists = p.clients[serviceType]; exists {
		return client, nil
	}

	client = p.createClient(serviceType, conn)
	p.clients[serviceType] = client

	p.logger.Info("new gRPC client created", "service", serviceType.String(), "address", addr)
	return client, nil
}
