package factory

import (
	"backend/protos/gen/go/sso"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type SSOClientType interface {
	sso.SSOServiceClient
	Close() error
}

type SSOClient struct {
	sso.SSOServiceClient
	conn *grpc.ClientConn
}

func (c *SSOClient) Close() error {
	return c.conn.Close()
}

func (p *ClientProvider) GetSSOClient(ctx context.Context) (SSOClientType, error) {
	client, err := p.getClient(ctx, ServiceSSO)
	if err != nil {
		return nil, err
	}
	ssoClient, ok := client.(SSOClientType)
	if !ok {
		return nil, fmt.Errorf("type assertion failed for SSO client")
	}
	return ssoClient, nil
}
