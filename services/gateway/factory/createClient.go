package factory

import (
	"backend/protos/gen/go/apps"
	"backend/protos/gen/go/production_task"
	"backend/protos/gen/go/sso"
	"google.golang.org/grpc"
)

func (p *ClientProvider) createClient(serviceType ServiceType, conn *grpc.ClientConn) interface{} {
	switch serviceType {
	case ServiceSSO:
		return &SSOClient{
			SSOServiceClient: sso.NewSSOServiceClient(conn),
			conn:             conn,
		}
	case ServiceApps:
		return &AppsClient{
			AppProviderServiceClient: app_provider.NewAppProviderServiceClient(conn),
			conn:                     conn,
		}
	case ServiceTasks:
		return &GetTasksClient{
			ProductionTaskServiceClient: production_task.NewProductionTaskServiceClient(conn),
			conn:                        conn,
		}
	default:
		panic("unknown service type")
	}
	return nil
}
