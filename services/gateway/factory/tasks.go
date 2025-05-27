package factory

import (
	"backend/protos/gen/go/production_task"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type TasksClientType interface {
	production_task.ProductionTaskServiceClient
	Close() error
}

type GetTasksClient struct {
	production_task.ProductionTaskServiceClient
	conn *grpc.ClientConn
}

func (c *GetTasksClient) Close() error {
	return c.conn.Close()
}

func (p *ClientProvider) GetTasksClient(ctx context.Context) (TasksClientType, error) {
	client, err := p.getClient(ctx, ServiceSSO)
	if err != nil {
		return nil, err
	}
	ssoClient, ok := client.(TasksClientType)
	if !ok {
		return nil, fmt.Errorf("type assertion failed for SSO client")
	}
	return ssoClient, nil
}
